package helm

import (
	"helm.sh/helm/v4/pkg/action"
	ri "helm.sh/helm/v4/pkg/release"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type Client struct {
	actRun           func(*action.List) ([]ri.Releaser, error)
	newList          func(*action.Configuration) *action.List
	restClientGetter genericclioptions.RESTClientGetter
}

func NewClient(restClientGetter genericclioptions.RESTClientGetter) *Client {
	return &Client{restClientGetter: restClientGetter, actRun: runAction, newList: action.NewList}
}

func (c *Client) ListReleases(namespace, name string, stateMask action.ListStates) ([]ri.Releaser, error) {
	helmCfg := &action.Configuration{}
	if err := helmCfg.Init(c.restClientGetter, namespace, ""); err != nil {
		return nil, err
	}
	l := c.newList(helmCfg)
	l.Filter = "^" + name + "$"
	l.StateMask = stateMask
	return c.actRun(l)
}

func runAction(l *action.List) ([]ri.Releaser, error) {
	return l.Run()
}
