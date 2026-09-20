package main 
import ("net/http"
"io"
"encoding/xml"
"context"
"html"
)

type RSSfeed struct {
 Channel struct {
			Title       string   `xml:"title"`
			Link        string        `xml:"link"`
			Description   string 	 `xml:"description"`
			Generator     string	  `xml:"generator"`
			Language       string		 `xml:"language"`
			LastBuildDate   string	 `xml:"lastBuildDate"`
			Item            []Item    `xml:"item"`
 } `xml:"channel"`
}

type Item struct {
	Title  string    `xml:"title"`
	Link   string    `xml:"link"`
	PubDate  string    `xml:"pubDate"`
	Guid      string  `xml:"guid"`
	Description  string `xml:"description"`
}


func fetchFeed(ctx context.Context, feedURL string) (*RSSfeed, error) {
	var feed RSSfeed;
	//create request
  req,err := http.NewRequestWithContext(ctx,"GET",feedURL,nil)

  if err != nil {
  return &feed,err
  }

req.Header.Set("User-Agent","gator")

client := &http.Client{}
res,err := client.Do(req)

if err != nil {
 return &feed,err
}
defer res.Body.Close()
bytes,err := io.ReadAll(res.Body)

if err != nil {
return &feed,err
}


if err := xml.Unmarshal(bytes,&feed);err != nil {
return &feed,err
}

feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
for _,i := range feed.Channel.Item {
	i.Title = html.UnescapeString(i.Title)
	i.Description = html.UnescapeString(i.Description)
}

return &feed,nil
  
}