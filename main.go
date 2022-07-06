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

var (
	srcCNFile string
	srcJPFile string
	srcUSFile string
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
)

func init()  {
	flag.StringVar(&srcCNFile, "scn", "ipip_cn.txt", "specify source ip list file")
	flag.StringVar(&srcJPFile, "sjp", "ipip_jp.txt", "specify source ip list file")
	flag.StringVar(&srcUSFile, "sus", "ipip_us.txt", "specify source ip list file")
	flag.StringVar(&dstFile, "d", "Country.mmdb", "specify destination mmdb file")
	flag.StringVar(&databaseType,"t", "GeoIP2-Country", "specify MaxMind database type")
	flag.Parse()
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

	var ipTxtListCN []string
	fh, err := os.Open(srcCNFile)
	if err != nil {
		log.Fatalf("fail to open %s\n", err)
	}
	scanner := bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipTxtListCN = append(ipTxtListCN, scanner.Text())
	}

	ipListJPCN := parseCIDRs(ipTxtListCN)
	for _, ip := range ipListJPCN {
		err = writer.Insert(ip, cnRecord)
		if err != nil {
			log.Fatalf("fail to insert to writer %v\n", err)
		}
	}

	var ipTxtListJP []string
	fh, err = os.Open(srcJPFile)
	if err != nil {
		log.Fatalf("fail to open %s\n", err)
	}
	scanner = bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipTxtListJP = append(ipTxtListJP, scanner.Text())
	}

	ipListJP := parseCIDRs(ipTxtListJP)
	for _, ip := range ipListJP {
		err = writer.Insert(ip, jpRecord)
		if err != nil {
			log.Fatalf("fail to insert to writer %v\n", err)
		}
	}

	var ipTxtListUS []string
	fh, err = os.Open(srcUSFile)
	if err != nil {
		log.Fatalf("fail to open %s\n", err)
	}
	scanner = bufio.NewScanner(fh)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipTxtListUS = append(ipTxtListUS, scanner.Text())
	}

	ipListUS := parseCIDRs(ipTxtListUS)
	for _, ip := range ipListUS {
		err = writer.Insert(ip, usRecord)
		if err != nil {
			log.Fatalf("fail to insert to writer %v\n", err)
		}
	}

	outFh, err := os.Create(dstFile)
	if err != nil {
		log.Fatalf("fail to create output file %v\n", err)
	}

	_, err = writer.WriteTo(outFh)
	if err != nil {
		log.Fatalf("fail to write to file %v\n", err)
	}

}


