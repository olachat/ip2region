package ip2region

//go:generate go run github.com/olachat/ip2region/v1.0/ip2location -csv ip2location/country.csv -out data/ip.ip2location.txt
//go:generate java -jar ./maker/java/dbMaker-1.2.2.jar -src ./data/ip.ip2location.txt -region ./data/global_region.csv
