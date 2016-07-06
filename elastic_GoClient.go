package main
import(
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"encoding/json"
)
func main(){
	url:="http://192.168.152.3:9200"
	client,err:=elastic.NewClient(elastic.SetURL(url))
	if err!=nil{
	        fmt.Println("Error occurred during Creating Client")		
	}

	info,code,err:=client.Ping(url).Do()
	if err!=nil{
		panic(err)
	}

	 fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	 fmt.Println("")


	esversion,err:=client.ElasticsearchVersion(url)
	if err!=nil{
		panic(err)
	}
	fmt.Println(esversion)
	
	exist,err:=client.IndexExists("chaoge1").Do()
	if err!=nil{
		 fmt.Println("Error occurred during judge exist")
	}
	println(exist)

	if !exist{
		createIndex,err:=client.CreateIndex("chaoge1").Do()
		if err!=nil{
			panic(err)
		}
		if !createIndex.Acknowledged{
			//
		}
	}
	type Tweet struct{
		User string
		Message string
		Retweets int	
	}
	tweet1:=Tweet{User:"chaoge",Message:"Give Me Five",Retweets:0}

	put2,err:=client.Index().
		Index("chaoge1").
		Type("test").
		Id("1").
		BodyJson(tweet1).
		Do()
	if err!=nil{

		panic(err)
	}
	tweet2:=Tweet{User:"chaoge",Message:"Give me six",Retweets:1}
	put3,err:=client.Index().
		Index("chaoge1").
		Type("test").
		Id("1").
		BodyJson(tweet2).
		Do()
	if err!=nil{
		panic(err)
	}
	tweet3:=Tweet{User:"chaoge",Message:"Give me seven",Retweets:2}
	put4,err:=client.Index().
		Index("chaoge1").
		Type("test").
		Id("1").
		BodyJson(tweet3).
		Do()
	if err!=nil{
		panic(err)
	}

	 fmt.Println("==============================")
	 fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	 fmt.Printf("Indexed tweet %s to index %s, type %s\n", put3.Id, put3.Index, put3.Type)
	 fmt.Printf("Indexed tweet %s to index %s, type %s\n", put4.Id, put4.Index, put4.Type)
	 fmt.Println("==============================")
	get1,err:=client.Get().
		Index("chaoge1").
		Type("test").
		Id("1").
		Do()
	if err!=nil{
		panic(err)
	}
	if get1.Found{
		 fmt.Println("")
		 fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		 fmt.Println("")
	}
	_,err=client.Flush().Index("chaoge1").Do()
	if err!=nil{
		panic(err)
	}
	
	termQuery:=elastic.NewTermQuery("User","chaoge")
	// 查询结果对Retweets进行汇总求和
	agg := elastic.NewSumAggregation().Field("Retweets")
	searchResult,err:=client.Search().
		Index("chaoge1").
		Query(termQuery).
		Aggregation("aggSum",agg).
		Sort("User",true).
		From(0).
		Size(10).
		Pretty(true).
		Do()
	if err!=nil{
		panic(err)
	}
	 fmt.Println("")			
	 fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	 fmt.Println("")
	
	if searchResult.Hits.TotalHits >0 {
		fmt.Printf("Found a total of %d tweets\n",searchResult.Hits.TotalHits)
		for _,hit :=range searchResult.Hits.Hits{
			var t Tweet
			err:=json.Unmarshal(*hit.Source,&t)
			if err!=nil{
				// Deserialization failed
			}
			fmt.Printf("Tweet by %s:%s and retweets is %d \n",t.User,t.Message,t.Retweets)
		}	
		
	}else{
		fmt.Print("found no  tweets\n")
	}
	agg1,found:=searchResult.Aggregations.Sum("aggSum")
	if !found{
		fmt.Println("Error occurred during found")
	}
		
	fmt.Printf("value:=============%v==============\n",int(*agg1.Value))				
	
}

