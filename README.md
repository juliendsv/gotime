gotime
======

Timestamp tool for humans.

##Installation


	go get github.com/juliendsv/gotime

Make sure the $GOPATH/bin is added into $PATH

##Usage

Hello World

	$ gotime -z yesterday 
	Local: 			 			2015-05-02 00:00:00 +0100 BST	
	UTC: 			 			2015-05-01 23:00:00 +0000 UTC
	timestamp: 		 			1430521200
	Milli timestamp: 			1430521200000
	Nano timestamp: 	 		1430521200000000000
	
Synopsis

	gotime [flags] [TIMESTAMP|EXPR]

##Example

	$ gotime
	Local: 			 			2015-05-03 12:55:54.646103581 +0100 BST
	UTC: 			 			2015-05-03 11:55:54.646103581 +0000 UTC
	timestamp: 		 			1430654154
	Milli timestamp: 	 		1430654154000
	Nano timestamp: 	 		1430654154646103581


	$ gotime 1349034753
	Local: 			 			2012-09-30 20:52:33 +0100 BST
	UTC: 			 			2012-09-30 19:52:33 +0000 UTC
	timestamp: 		 			1349034753
	Milli timestamp: 	 		1349034753000
	Nano timestamp: 	 		1349034753000000000


	$ gotime yesterday
	Local: 			 			2015-05-02 12:55:54.646459355 +0100 BST
	UTC: 			 			2015-05-02 11:55:54.646459355 +0000 UTC
	timestamp: 		 			1430567754
	Milli timestamp: 	 		1430567754000
	Nano timestamp: 	 		1430567754646459355


	$ gotime -z yesterday
	Local: 			 			2015-05-02 00:00:00 +0100 BST
	UTC: 			 			2015-05-01 23:00:00 +0000 UTC
	timestamp: 		 			1430521200
	Milli timestamp: 	 		1430521200000
	Nano timestamp: 	 		1430521200000000000


	$ gotime tomorrow
	Local: 						2015-05-04 12:55:54.646514439 +0100 BST
	UTC: 			 			2015-05-04 11:55:54.646514439 +0000 UTC
	timestamp: 		 			1430740554
	Milli timestamp: 	 		1430740554000
	Nano timestamp: 	 		1430740554646514439


	$ gotime 5daysago
	Local: 			 			2015-04-28 12:55:54.646535822 +0100 BST
	UTC: 			 			2015-04-28 11:55:54.646535822 +0000 UTC
	timestamp: 		 			1430222154
	Milli timestamp: 	 		1430222154000
	Nano timestamp: 	 		1430222154646535822


	$ gotime 5dayfwd
	Local: 			 			2015-05-08 12:55:54.646704671 +0100 BST
	UTC: 			 			2015-05-08 11:55:54.646704671 +0000 UTC
	timestamp: 		 			1431086154
	Milli timestamp: 	 		1431086154000
	Nano timestamp: 		 	1431086154646704671


	$ gotime 1405967017972502579
	Local: 			 			2014-07-21 19:23:37.972502579 +0100 BST
	UTC: 			 			2014-07-21 18:23:37.972502579 +0000 UTC
	timestamp: 		 			1405967017
	Milli timestamp: 	 		1405967017000
	Nano timestamp: 	 		1405967017972502579


	$ gotime "2014-02-03 19:54:02"
	Local: 			 			2014-02-03 19:54:02 +0000 GMT
	UTC: 			 			2014-02-03 19:54:02 +0000 UTC
	timestamp: 		 			1391457242
	Milli timestamp: 	 		1391457242000
	Nano timestamp: 	 		1391457242000000000

