package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/syyongx/php2go"
)

var CountryNames = map[string]string{
	"AD": "安道尔",
	"AE": "阿联酋",
	"AF": "阿富汗",
	"AG": "安提瓜和巴布达",
	"AI": "安圭拉",
	"AL": "阿尔巴尼亚",
	"AM": "亚美尼亚",
	"AO": "安哥拉",
	"AR": "阿根廷",
	"AS": "美属萨摩亚",
	"AT": "奥地利",
	"AU": "澳大利亚",
	"AW": "阿鲁巴",
	"AZ": "阿塞拜疆",
	"BA": "波黑",
	"BB": "巴巴多斯",
	"BD": "孟加拉",
	"BE": "比利时",
	"BF": "布基纳法索",
	"BG": "保加利亚",
	"BH": "巴林",
	"BI": "布隆迪",
	"BJ": "贝宁",
	"BM": "百慕大",
	"BN": "文莱",
	"BO": "玻利维亚",
	"BR": "巴西",
	"BS": "巴哈马",
	"BT": "不丹",
	"BW": "博茨瓦纳",
	"BY": "白俄罗斯",
	"BZ": "伯利兹",
	"CA": "加拿大",
	"CD": "刚果金",
	"CG": "刚果布",
	"CH": "瑞士",
	"CI": "科特迪瓦",
	"CL": "智利",
	"CM": "喀麦隆",
	"CN": "中国",
	"CO": "哥伦比亚",
	"CR": "哥斯达黎加",
	"CV": "佛得角",
	"CW": "库拉索",
	"CY": "塞浦路斯",
	"CZ": "捷克",
	"DE": "德国",
	"DJ": "吉布提",
	"DK": "丹麦",
	"DM": "多米尼克",
	"DO": "多米尼加",
	"DZ": "阿尔及利亚",
	"EC": "厄瓜多尔",
	"EE": "爱沙尼亚",
	"EG": "埃及",
	"ES": "西班牙",
	"ET": "埃塞俄比亚",
	"EU": "欧洲",
	"FI": "芬兰",
	"FJ": "斐济",
	"FM": "密克罗尼西亚",
	"FR": "法国",
	"GA": "加蓬",
	"GB": "英国",
	"GD": "格林纳达",
	"GE": "格鲁吉亚",
	"GH": "加纳",
	"GI": "直布罗陀",
	"GM": "冈比亚",
	"GN": "几内亚",
	"GQ": "赤道几内亚",
	"GR": "希腊",
	"GT": "危地马拉",
	"GU": "关岛",
	"GW": "几内亚比绍",
	"GY": "圭亚那",
	"HK": "香港",
	"HN": "洪都拉斯",
	"HR": "克罗地亚",
	"HT": "海地",
	"HU": "匈牙利",
	"ID": "印度尼西亚",
	"IE": "爱尔兰",
	"IL": "以色列",
	"IN": "印度",
	"IQ": "伊拉克",
	"IR": "伊朗",
	"IS": "冰岛",
	"IT": "意大利",
	"JM": "牙买加",
	"JO": "约旦",
	"JP": "日本",
	"KE": "肯尼亚",
	"KG": "吉尔吉斯斯坦",
	"KH": "柬埔寨",
	"KI": "基里巴斯",
	"KM": "科摩罗",
	"KN": "圣基茨和尼维斯",
	"KP": "朝鲜",
	"KR": "韩国",
	"KW": "科威特",
	"KY": "开曼群岛",
	"KZ": "哈萨克斯坦",
	"LA": "老挝",
	"LB": "黎巴嫩",
	"LC": "圣卢西亚",
	"LI": "列支敦士登",
	"LK": "斯里兰卡",
	"LR": "利比里亚",
	"LS": "莱索托",
	"LT": "立陶宛",
	"LU": "卢森堡",
	"LV": "拉脱维亚",
	"LY": "利比亚",
	"MA": "摩洛哥",
	"MC": "摩纳哥",
	"MD": "摩尔多瓦",
	"ME": "黑山",
	"MG": "马达加斯加",
	"MK": "马其顿",
	"ML": "马里",
	"MM": "缅甸",
	"MN": "蒙古",
	"MO": "澳门",
	"MR": "毛里塔尼亚",
	"MS": "蒙塞拉特岛",
	"MT": "马耳他",
	"MU": "毛里求斯",
	"MV": "马尔代夫",
	"MW": "马拉维",
	"MX": "墨西哥",
	"MY": "马来西亚",
	"MZ": "莫桑比克",
	"NA": "纳米比亚",
	"NE": "尼日尔",
	"NG": "尼日利亚",
	"NI": "尼加拉瓜",
	"NL": "荷兰",
	"NO": "挪威",
	"NP": "尼泊尔",
	"NU": "纽埃",
	"NZ": "新西兰",
	"OM": "阿曼",
	"PA": "巴拿马",
	"PE": "秘鲁",
	"PG": "巴布亚新几内亚",
	"PH": "菲律宾",
	"PK": "巴基斯坦",
	"PL": "波兰",
	"PR": "波多黎各",
	"PS": "巴勒斯坦",
	"PT": "葡萄牙",
	"PY": "巴拉圭",
	"QA": "卡塔尔",
	"RO": "罗马尼亚",
	"RS": "塞尔维亚",
	"RU": "俄罗斯",
	"RW": "卢旺达",
	"SA": "沙特阿拉伯",
	"SB": "所罗门群岛",
	"SC": "塞舌尔",
	"SE": "瑞典",
	"SG": "新加坡",
	"SI": "斯洛文尼亚",
	"SK": "斯洛伐克",
	"SL": "塞拉利昂",
	"SM": "圣马力诺",
	"SN": "塞内加尔",
	"SR": "苏里南",
	"SV": "萨尔瓦多",
	"SY": "叙利亚",
	"SZ": "斯威士兰",
	"TC": "特克斯和凯科斯群岛",
	"TD": "乍得",
	"TG": "多哥",
	"TH": "泰国",
	"TJ": "塔吉克斯坦",
	"TM": "土库曼斯坦",
	"TN": "突尼斯",
	"TO": "汤加",
	"TR": "土耳其",
	"TT": "特立尼达和多巴哥",
	"TW": "台湾",
	"TZ": "坦桑尼亚",
	"UA": "乌克兰",
	"UG": "乌干达",
	"US": "美国",
	"UY": "乌拉圭",
	"UZ": "乌兹别克斯坦",
	"VA": "梵蒂冈",
	"VC": "圣文森特和格林纳丁斯",
	"VE": "委内瑞拉",
	"VG": "英属维尔京群岛",
	"VI": "美属维尔京群岛",
	"VN": "越南",
	"VU": "瓦努阿图",
	"WS": "萨摩亚",
	"YE": "也门",
	"ZA": "南非",
	"ZM": "赞比亚",
	"ZW": "津巴布韦",
}

func main() {
	// Parse command-line arguments
	csvPath := flag.String("csv", "", "path to CSV file")
	outPath := flag.String("out", "", "path to output file")
	flag.Parse()

	// Make sure a CSV path was provided
	if *csvPath == "" {
		fmt.Println("Error: Please provide a path to a CSV file using the -csv flag")
		return
	}

	// Open the CSV file
	file, err := os.Open(*csvPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Open the output file
	outFile, err := os.Create(*outPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	// Create a CSV reader and read the file line by line
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading record:", err)
			continue
		}

		// Parse the start and end IP addresses
		startInt := parseIP(record[0])
		endInt := parseIP(record[1])

		startIp := php2go.Long2ip(uint32(startInt))
		endIp := php2go.Long2ip(uint32(endInt))

		country, has := CountryNames[record[2]]
		if has {
			outFile.WriteString(fmt.Sprintf("%s|%s|%s|0|0|0|0\n", startIp, endIp, country))
		} else {
			fmt.Println("Error: Unknown country code:", record[2], record[3])
		}
	}
}

func parseIP(ipLong string) int {
	startInt, err := strconv.Atoi(ipLong)
	if err != nil {
		panic(fmt.Sprintln("Error parsing start IP address", ipLong, "err:", err))
	}
	return startInt
}
