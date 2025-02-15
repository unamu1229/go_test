<?php


class ElasticSearch
{

    public function get()
    {
        $ch = curl_init("http://elasticsearch:9200/search_job/_search");
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

    public function make(int $id)
    {
        $ch = curl_init("http://elasticsearch:9200/search_job/_doc/" . $id);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "PUT");
        curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($this->station($id)));
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json;  charset=UTF-8']);
        curl_exec($ch);
        curl_close($ch);
    }

    public function station(int $id)
    {
        return [
            'pref' => '東京' . $id,
            'employment' => 'full-time',
            'line' => [
                'name' => '山手線',
                'station' => [
                    '東京'
                ]
            ]
        ];
    }
}
echo 'start';

if (!$hoge = 0) {
    var_dump($hoge);
}

var_dump($hoge);

return;

$start = microtime(true);
$es = new ElasticSearch();
$es->get();
//$es->make(2);
//for ($i = 0; $i < 100; $i++) {
//    $es->make($i);
//}
var_dump(microtime(true) - $start);