/*
 * @Author: tqcenglish
 * @LastEditors: tqcenglish
 * @Email: tqcenglish#gmail.com
 * @Description: 一梦如是，总归虚无
 * @LastEditTime: 2019-04-15 17:14:33
 */

package routes

import "k-peach/pkg/context"

//Pong 响应 ping
func Pong(ctx *context.Context) {
	ctx.JSON(200, map[string]string{"message": "pong"})
}
