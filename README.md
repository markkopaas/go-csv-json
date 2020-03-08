Start redis:
    
    docker run -d -p 6379:6379 --name redis1 redis
    
Process csv file with template, and send output to redis
    
    go run main.go -template "$(< redis.template)" <example.csv  | nc localhost 6379
    
Test read:
    
    echo "GET 41.5" | nc localhost 6379
    
    # {"City":"You\"ng'sto\nwn","EW":"W","LatD":"41","LatM":"5","LatS":"59","LonD":"80","LonM":"39","LonS":"0","NS":"N","State":"OH"}
    
Big test:

    $ source generatecsv.sh 30000 >30k.csv
    $ source generatecsv.sh 3000000 >3m.csv
    $ tail 3m.csv
    $ time (go run main.go -template "$(< redis.template)" <3m.csv  | nc localhost 6379 > /dev/null)
    
    real	2m29.797s
    user	1m28.762s
    sys	1m8.159s
    
Inspect the store:

    $ echo "DBSIZE" | nc localhost 6379
    $ echo "GET 3000000.52" | nc localhost 6379
