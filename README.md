#Steam Crawler
[![GoDoc](https://godoc.org/github.com/Rompei/steam-crawler?status.png)](https://godoc.org/github.com/Rompei/steam-crawler)
[![Build Status](https://drone.io/github.com/Rompei/steam-crawler/status.png)](https://drone.io/github.com/Rompei/steam-crawler/latest)

Steam Crawler is crawler for discounted games in Steam.  


##Usage

```
go get github.com/Rompei/steam-crawler

If you got error, "socket: too many open files", please set `ulimit -n` more, because this software opens sockets more than ulimit already set.
```

##Data Structure in CSV

```
Nnumber,Name,Release date,Discount rate,Normal price,Discount price,Reputation,the number of reviewer,URL
```

###Example

```
1,Counter-Strike: Global Offensive,"21 Aug, 2012",50,1480,740,93,948815,http://store.steampowered.com/app/730/
```

##License

[BSD 3-Clause license](http://opensource.org/licenses/BSD-3-Clause)

##Note

Please use this software by self-responsibility.
