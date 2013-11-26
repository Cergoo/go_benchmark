<?php

class bench {

    private static $ArrayLabel;

    public static function start($Label) {
        self::$ArrayLabel[$Label] = microtime(true);
    }

    public static function stop($Label) {
        $delta = round(microtime(true) - self::$ArrayLabel[$Label], 4);
        echo 'checkpoint_', $Label, '<br>time : ', $delta, ' сек.';
    }

}

$v = [
    'Itm_1' => 100,
    'Itm_2' => 'aaaaaaaaaaaaaaaabbbbbbbbbbbbbb',
    'Itm_3' => true,
    'Itm_4' => 17.257,
    'Itm_5' => 117,
    'Itm_6' => ["aaaaaaaaaaaaaaaabbbbbbbbbbbbbb1", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb2", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb3", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb4", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb5", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb6", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb7", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb8", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb9", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb10"],
    'Itm_7' => [1, 2, 4, 7, 8, 9, 11, 145, 18, 20],
    'Itm_8' => [true, false, true, false, true, false, true, false, true, false],
    'Itm_9' => [11.2, 12.1, 17.1, 11.1, 14.21],
    'Itm10' => ['key1'  => 'itm1', 'key2'  => 'itm2', 'key3'  => 'itm3', 'key4'  => 'itm4', 'key5'  => 'itm5', 'key6'  => 'itm6', 'key7'  => 'itm7', 'key8'  => 'itm8', 'key9'  => 'itm9', 'key10' => 'itm10',],
    'Itm11' => ['key1'  => 1, 'key2'  => 2, 'key3'  => 10, 'key4'  => 12, 'key5'  => 12, 'key6'  => 14, 'key7'  => 12, 'key8'  => 12, 'key9'  => 12, 'key10' => 14,],
    'Itm12' => ['key1'  => true, 'key2'  => false, 'key3'  => true, 'key4'  => false, 'key5'  => true, 'key6'  => true, 'key7'  => false, 'key8'  => true, 'key9'  => false, 'key10' => true,],
    'Itm13' => [],
];

$v1         = $v;
$v['Itm13'] = &$v1;

bench::start('1');
for ($x = 0; $x < 1000; $x++) {
    $n = json_encode($v);
}
bench::stop('1');

print_r($n);
    ?>