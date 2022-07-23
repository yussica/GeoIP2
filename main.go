package main

import (
	"bufio"
	"flag"
	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	log "github.com/sirupsen/logrus"
	"os"
)
		//"country": mmdbtype.Map{
		//"continent": mmdbtype.Map{
		//"code": mmdbtype.String("AS"),
		//"geoname_id": mmdbtype.Uint32(6255147),
		//"names":mmdbtype.Map{"de":mmdbtype.String("Asien"),
		//"en":mmdbtype.String("Asia"),
		//"es":mmdbtype.String("Asia"),
		//"fr":mmdbtype.String("Asie"),
		//"ja":mmdbtype.String("アジア"),
		//"pt-BR":mmdbtype.String("Ásia"),
		//"ru":mmdbtype.String("Азия"),
		//"zh-CN":mmdbtype.String("亚洲")},
		//},
		//"country": mmdbtype.Map{
		//"geoname_id":mmdbtype.Uint32(1814991),
		//"is_in_european_union":mmdbtype.Bool(false),
		//"iso_code":mmdbtype.String("CN"),
		//"names":mmdbtype.Map{
		//"de":mmdbtype.String("China"),
		//"en":mmdbtype.String("China"),
		//"es":mmdbtype.String("China"),
		//"fr":mmdbtype.String("Chine"),
		//"ja":mmdbtype.String("中国"),
		//"pt-BR":mmdbtype.String("China"),
		//"ru":mmdbtype.String("Китай"),
		//"zh-CN":mmdbtype.String("中国"),
		//},
		//},
		//"registered_country": mmdbtype.Map{
		//"geoname_id":mmdbtype.Uint32(1814991),
		//"is_in_european_union":mmdbtype.Bool(false),
		//"iso_code":mmdbtype.String("CN"),
		//"names":mmdbtype.Map{
		//"de":mmdbtype.String("China"),
		//"en":mmdbtype.String("China"),
		//"es":mmdbtype.String("China"),
		//"fr":mmdbtype.String("Chine"),
		//"ja":mmdbtype.String("中国"),
		//"pt-BR":mmdbtype.String("China"),
		//"ru":mmdbtype.String("Китай"),
		//"zh-CN":mmdbtype.String("中国"),
		//},
		//},
		//"traits": mmdbtype.Map{
		//"is_anonymous_proxy": mmdbtype.Bool(false),
		//"is_satellite_provider":mmdbtype.Bool(false),
		//},
		//},
//http://www.geonames.org/
var (
	srcCNFile string
	srcJPFile string
	srcUSFile string
	srcHKFile string
	dstFile string
	databaseType string
	cnRecord = mmdbtype.Map{
		"country": mmdbtype.Map{
			"geoname_id":           mmdbtype.Uint32(1814991),
			"is_in_european_union": mmdbtype.Bool(false),
			"iso_code":             mmdbtype.String("CN"),
			"names": mmdbtype.Map{
				"de":    mmdbtype.String("China"),
				"en":    mmdbtype.String("China"),
				"es":    mmdbtype.String("China"),
				"fr":    mmdbtype.String("Chine"),
				"ja":    mmdbtype.String("中国"),
				"pt-BR": mmdbtype.String("China"),
				"ru":    mmdbtype.String("Китай"),
				"zh-CN": mmdbtype.String("中国"),
			},
		},
	}
	jpRecord = mmdbtype.Map{
		"country": mmdbtype.Map{
			"geoname_id":           mmdbtype.Uint32(1861060),
			"is_in_european_union": mmdbtype.Bool(false),
			"iso_code":             mmdbtype.String("JP"),
			"names": mmdbtype.Map{
				"de":    mmdbtype.String("Japan"),
				"en":    mmdbtype.String("Japan"),
				"es":    mmdbtype.String("Japón"),
				"fr":    mmdbtype.String("Japon"),
				"ja":    mmdbtype.String("日本"),
				"pt-BR": mmdbtype.String("Japão"),
				"ru":    mmdbtype.String("Япония"),
				"zh-CN": mmdbtype.String("日本"),
			},
		},
	}
	usRecord = mmdbtype.Map{
		"country": mmdbtype.Map{
			"geoname_id":           mmdbtype.Uint32(6252001),
			"is_in_european_union": mmdbtype.Bool(false),
			"iso_code":             mmdbtype.String("US"),
			"names": mmdbtype.Map{
				"de":    mmdbtype.String("vereinigte Staaten von Amerika"),
				"en":    mmdbtype.String("United States"),
				"es":    mmdbtype.String("Estados Unidos de America"),
				"fr":    mmdbtype.String("les États-Unis d'Amérique"),
				"ja":    mmdbtype.String("アメリカ"),
				"pt-BR": mmdbtype.String("Estados Unidos da América"),
				"ru":    mmdbtype.String("Соединенные Штаты Америки"),
				"zh-CN": mmdbtype.String("美国"),
			},
		},
	}
	hkRecord = mmdbtype.Map{
		"country": mmdbtype.Map{
			"geoname_id":           mmdbtype.Uint32(1819729),
			"is_in_european_union": mmdbtype.Bool(false),
			"iso_code":             mmdbtype.String("HK"),
			"names": mmdbtype.Map{
				"de":    mmdbtype.String("Hongkong"),
				"en":    mmdbtype.String("Hong Kong"),
				"es":    mmdbtype.String("Hong Kong"),
				"fr":    mmdbtype.String("Hong Kong"),
				"ja":    mmdbtype.String("香港"),
				"pt-BR": mmdbtype.String("Hong Kong"),
				"ru":    mmdbtype.String("Гонконг"),
				"zh-CN": mmdbtype.String("香港"),
			},
		},
	}
)

func init()  {
	flag.StringVar(&srcCNFile, "scn", "ipip_cn.txt", "specify source ip list file")
	flag.StringVar(&srcJPFile, "sjp", "ipip_jp.txt", "specify source ip list file")
	flag.StringVar(&srcUSFile, "sus", "ipip_us.txt", "specify source ip list file")
	flag.StringVar(&srcHKFile, "shk", "ipip_hk.txt", "specify source ip list file")
	flag.StringVar(&dstFile, "d", "Country.mmdb", "specify destination mmdb file")
	flag.StringVar(&databaseType,"t", "GeoIP2-Country", "specify MaxMind database type")
	flag.Parse()
}

func addGEOIP(writer *mmdbwriter.Tree,srcFile string, record mmdbtype.Map) {
	var ipTxtList []string
	fh, err := os.Open(srcFile)
	if err != nil {
		log.Fatalf("fail to open %s\n", err)
	}
	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipTxtList = append(ipTxtList, scanner.Text())
	}

	log.Printf("%v start to parse to CIDR\n", srcFile)
	ipListHK := parseCIDRs(ipTxtList)
	for _, ip := range ipListHK {
		err = writer.Insert(ip, record)
		if err != nil {
			log.Fatalf("fail to insert to writer %v\n", err)
		}
	}
}

func main()  {
	writer, err := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType: databaseType,
			RecordSize:   24,
		},
	)
	if err != nil {
		log.Fatalf("fail to new writer %v\n", err)
	}

	addGEOIP(writer, srcCNFile, cnRecord)
	addGEOIP(writer, srcJPFile, jpRecord)
	addGEOIP(writer, srcUSFile, usRecord)
	addGEOIP(writer, srcHKFile, hkRecord)

	outFh, err := os.Create(dstFile)
	if err != nil {
		log.Fatalf("fail to create output file %v\n", err)
	}

	_, err = writer.WriteTo(outFh)
	if err != nil {
		log.Fatalf("fail to write to file %v\n", err)
	}

}


