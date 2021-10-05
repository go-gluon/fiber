package fiber

import config "github.com/go-gluon/gluon/config"

func (item *FiberConfig) ReadFromMapNode(node config.MapNode) error {
	item.Enabled = node.Bool("enabled", item.Enabled)
	item.Listen = node.String("listen", item.Listen)
	item.DisableStartupMessage = node.Bool("disable-startup-message", item.DisableStartupMessage)
	return nil
}
