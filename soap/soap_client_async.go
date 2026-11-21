package melipayamak

import "context"

func (c *SOAPClient) SendSimpleSMS2Async(ctx context.Context, to, from, text string, flash bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendSimpleSMS2(ctx, to, from, text, flash) })
}
