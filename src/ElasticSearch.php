<?php


class ElasticSearch
{

    public function get()
    {
        $ch = curl_init("http://172.19.0.3:9200/search_job/_search");
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "GET");
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
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