package wechat

import (
	"errors"
	"github.com/chanxuehong/wechat/menu"
)

// 创建自定义菜单.
//  NOTE: 创建自定义菜单后，由于微信客户端缓存，需要24小时微信客户端才会展现出来。
//  建议测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。
func (c *Client) MenuCreate(_menu *menu.Menu) (err error) {
	if _menu == nil {
		return errors.New("_menu == nil")
	}

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := clientMenuCreateURL(token)

	var result Error
	if err = c.postJSON(_url, _menu, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return &result
	}
	return
}

// 删除自定义菜单
func (c *Client) MenuDelete() error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := clientMenuDeleteURL(token)

	var result Error
	if err = c.getJSON(_url, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

// 获取自定义菜单
func (c *Client) MenuGet() (*menu.Menu, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMenuGetURL(token)

	var result struct {
		Menu menu.Menu `json:"menu"`
		Error
	}
	if err = c.getJSON(_url, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.Menu, nil
}
