package ippool

var ip_pool map[int]string

func init() {
	ip_pool = make(map[int]string)

	ip_pool[0] = "http://112.87.68.168:9999"
	ip_pool[1] = "http://112.85.166.133:9999"
	ip_pool[2] = "http://111.202.37.195:42558"
	ip_pool[3] = "http://125.123.138.70:9000"
	ip_pool[4] = "http://119.187.120.118:8060"
	ip_pool[5] = "http://119.145.2.99:44129"
	ip_pool[6] = "http://180.107.95.163:8118"
	ip_pool[7] = "http://103.235.199.93:31854"
	ip_pool[8] = "http://163.204.247.206:9999"
	ip_pool[9] = "http://218.91.112.108:9999"
}

func GetIP() string {
	return ""
	//index := rand.Intn(10)
	//
	//if index  == 10 {
	//	return ""
	//}
	//
	//ip := ip_pool[index]
	//
	//return ip
}
