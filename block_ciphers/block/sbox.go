package block

type Sbox [8][16]uint8

var TestKey = []uint8{0xbe, 0x5e, 0xc2, 0x00, 0x6c, 0xff, 0x9d, 0xcf,
	0x52, 0x35, 0x49, 0x59, 0xf1, 0xff, 0x0c, 0xbf,
	0xe9, 0x50, 0x61, 0xb5, 0xa6, 0x48, 0xc1, 0x03,
	0x87, 0x06, 0x9c, 0x25, 0x99, 0x7c, 0x06, 0x72}

var TestSBox = Sbox([8][16]uint8{
	{4, 10, 9, 2, 13, 8, 0, 14, 6, 11, 1, 12, 7, 15, 5, 3},
	{14, 11, 4, 12, 6, 13, 15, 10, 2, 3, 8, 1, 0, 7, 5, 9},
	{5, 8, 1, 13, 10, 3, 4, 2, 14, 15, 12, 7, 6, 0, 9, 11},
	{7, 13, 10, 1, 0, 8, 9, 15, 14, 4, 6, 12, 11, 2, 5, 3},
	{6, 12, 7, 1, 5, 15, 13, 8, 4, 10, 9, 14, 0, 3, 11, 2},
	{4, 11, 10, 0, 7, 2, 1, 13, 3, 6, 8, 5, 9, 12, 15, 14},
	{13, 11, 4, 1, 3, 15, 5, 9, 0, 10, 14, 7, 6, 8, 2, 12},
	{1, 15, 13, 0, 5, 7, 10, 4, 9, 2, 3, 14, 6, 11, 8, 12},
})

func (s *Sbox) k(n nv) nv {
	return nv(s[0][(n>>0)&0x0F])<<0 +
		nv(s[1][(n>>4)&0x0F])<<4 +
		nv(s[2][(n>>8)&0x0F])<<8 +
		nv(s[3][(n>>12)&0x0F])<<12 + 
		nv(s[4][(n>>16)&0x0F])<<16 +
		nv(s[5][(n>>20)&0x0F])<<20 +
		nv(s[6][(n>>24)&0x0F])<<24 +
		nv(s[7][(n>>28)&0x0F])<<28
}