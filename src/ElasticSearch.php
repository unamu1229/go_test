<?php


class ElasticSearch
{

    public function get()
    {
        $ch = curl_init("http://172.54.0.54:9200/search_job/_search");
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "GET");
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);

        for ($i = 0; $i < 100; $i++ ) {
            echo $i;
            curl_exec($ch);
        }
        $response = curl_exec($ch);
        var_dump($response);

        curl_close($ch);
    }
}
echo 'start';
$start = microtime(true);
$es = new ElasticSearch();
$es->get();
var_dump(microtime(true) - $start);