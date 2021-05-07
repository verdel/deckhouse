package checker

import (
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"d8.io/upmeter/pkg/check"
	k8s "d8.io/upmeter/pkg/kubernetes"
	"d8.io/upmeter/pkg/probe/util"
)

// NamespaceLifecycle is a checker constructor and configurator
type NamespaceLifecycle struct {
	Access                   *k8s.Access
	CreationTimeout          time.Duration
	DeletionTimeout          time.Duration
	GarbageCollectionTimeout time.Duration
}

func (c NamespaceLifecycle) Checker() check.Checker {
	return &namespaceLifeCycleChecker{
		access:                  c.Access,
		creationTimeout:         c.CreationTimeout,
		deletionTimeout:         c.DeletionTimeout,
		garbageCollectorTimeout: c.GarbageCollectionTimeout,
	}
}

type namespaceLifeCycleChecker struct {
	access          *k8s.Access
	creationTimeout time.Duration
	deletionTimeout time.Duration

	garbageCollectorTimeout time.Duration

	// inner state
	checker check.Checker
}

func (c *namespaceLifeCycleChecker) BusyWith() string {
	return c.checker.BusyWith()
}

func (c *namespaceLifeCycleChecker) Check() check.Error {
	namespace := createNamespaceObject()
	c.checker = c.new(namespace)
	return c.checker.Check()
}

/*
1. check control plane availability
2. collect the garbage of the namespace from previous runs
3. create and delete the namespace in api
4. ensure it does not exist (with retries)
*/
func (c *namespaceLifeCycleChecker) new(namespace *v1.Namespace) check.Checker {
	kind := namespace.GetObjectKind().GroupVersionKind().Kind
	name := namespace.GetName()
	labels := namespace.GetLabels()

	notListedChecker := withRetryEachSeconds(
		&objectIsNotListedChecker{
			access:   c.access,
			kind:     kind,
			listOpts: listOptsByName(name),
		},
		c.deletionTimeout)

	check := sequence(
		&controlPlaneChecker{c.access},
		newGarbageCollectorCheckerByLabels(c.access, kind, "", labels, c.garbageCollectorTimeout),
		withTimeout(&namespaceCreationChecker{access: c.access, namespace: namespace}, c.creationTimeout),
		withTimeout(&namespaceDeletionChecker{access: c.access, namespace: namespace}, c.deletionTimeout),
		notListedChecker,
	)

	return withTimeout(check, c.deletionTimeout)
}

// namespaceCreationChecker creates namespace
type namespaceCreationChecker struct {
	access    *k8s.Access
	namespace *v1.Namespace
}

func (c *namespaceCreationChecker) BusyWith() string {
	return fmt.Sprintf("creating namespace %q", c.namespace.GetName())
}

func (c *namespaceCreationChecker) Check() check.Error {
	client := c.access.Kubernetes()
	_, err := client.CoreV1().Namespaces().Create(c.namespace)
	if err != nil {
		return check.ErrUnknown("cannot create namespace %q: %v", c.namespace.GetName(), err)
	}
	return nil
}

// namespaceDeletionChecker deletes namespace
type namespaceDeletionChecker struct {
	access    *k8s.Access
	namespace *v1.Namespace
}

func (c *namespaceDeletionChecker) BusyWith() string {
	return fmt.Sprintf("deleting namespace %q", c.namespace.GetName())
}

func (c *namespaceDeletionChecker) Check() check.Error {
	client := c.access.Kubernetes()
	err := client.CoreV1().Namespaces().Delete(c.namespace.GetName(), &metav1.DeleteOptions{})
	if err != nil {
		return check.ErrFail("cannot delete namespace %q: %v", c.namespace.GetName(), err)
	}
	return nil
}

func createNamespaceObject() *v1.Namespace {
	name := util.RandomIdentifier("upmeter-control-plane-namespace")

	return &v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"heritage":      "upmeter",
				"upmeter-agent": util.AgentUniqueId(),
				"upmeter-group": "control-plane",
				"upmeter-probe": "namespace",
			},
		},
	}
}
