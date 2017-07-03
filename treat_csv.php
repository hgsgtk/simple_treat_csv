<?php
	/*
	  define database connection
	 */
	 $dbname = "treat_csv";
	 $socket_path = '/Applications/MAMP/tmp/mysql/mysql.sock';

	 $dsn = "mysql:dbname={$dbname};host={$dbhost};unix_socket={$socket_path}";
	 $dbuser = "root";
	 $dbpassword = "root";

	/*
	  define csv file information
	*/
	$file_path = "customer.csv";
	$export_csv_title = ["顧客ID", "名前", "性別", "電話番号"];
	$export_sql = "SELECT id, name, sex, tel FROM customers";

	/*
	  Create Database Connection
	 */
	try{
		$dbh = new PDO($dsn, $dbuser, $dbpassword);
	}catch(PDOException $e){
		print('Connection failed:'.$e->getMessage());
		die();
	}

 	// encoding title into SJIS-win
    foreach( $export_csv_title as $key => $val ){
             
        $export_header[] = mb_convert_encoding($val, 'SJIS-win', 'UTF-8');
    }
	/*
		Make CSV content
	 */
	if(touch($file_path)){
		$file = new SplFileObject($file_path, "w");

		// write csv header
		$file->fputcsv($export_header);

		// query database
		$stmt = $dbh->query($export_sql);

		// create csv sentences
		$csv_str = "";
		while($row = $stmt->fetch(PDO::FETCH_ASSOC)){
			$file->fputcsv($row);

		}

		// close database connection
		$dbh = null;
	}