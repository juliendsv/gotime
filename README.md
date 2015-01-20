gotime
======

Give it a timestamp or a nano-timestamp it will give you the time in a readable format.

##Installation


	go get github.com/juliendsv/gotime


##Usage


	$ gotime
	-----------------
	UTC: 			 	 	2015-01-20 17:20:13.349027819 +0000 UTC
	Local: 			 	 	2015-01-20 17:20:13.349027819 +0000 GMT
	timestamp: 		 	 	1421774413
	Milli timestamp: 	 	1421774413000
	Nano timestamp: 	 	1421774413349027819

	$ gotime 1349034753
	-----------------
	UTC: 			 		2012-09-30 19:52:33 +0000 UTC
	Local: 			 		2012-09-30 20:52:33 +0100 BST
	timestamp: 		 		1349034753
	Milli timestamp: 	 	1349034753000
	Nano timestamp: 	 	1349034753000000000

	$ gotime yesterday
	-----------------
	UTC: 			 		2015-01-19 17:20:13.349240058 +0000 UTC
	Local: 			 		2015-01-19 17:20:13.349240058 +0000 GMT
	timestamp: 		 		1421688013
	Milli timestamp: 	 	1421688013000
	Nano timestamp: 	 	1421688013349240058

	$ gotime -z yesterday
	-----------------
	UTC: 			 		2015-01-19 00:00:00 +0000 UTC
	Local: 			 		2015-01-19 00:00:00 +0000 GMT
	timestamp: 		 		1421625600
	Milli timestamp: 	 	1421625600000
	Nano timestamp: 	 	1421625600000000000

	$ gotime tomorrow
	-----------------
	UTC: 			 		2015-01-21 17:20:13.349284741 +0000 UTC
	Local: 			 		2015-01-21 17:20:13.349284741 +0000 GMT
	timestamp: 		 		1421860813
	Milli timestamp: 	 	1421860813000
	Nano timestamp: 	 	1421860813349284741

	$ gotime 5daysago
	-----------------
	UTC: 			 		2015-01-15 17:20:13.349464529 +0000 UTC
	Local: 			 		2015-01-15 17:20:13.349464529 +0000 GMT
	timestamp: 		 		1421342413
	Milli timestamp: 	 	1421342413000
	Nano timestamp: 	 	1421342413349464529

	$ gotime 1405967017972502579
	-----------------
	UTC: 			 		2014-07-21 18:23:37.972502579 +0000 UTC
	Local: 			 		2014-07-21 19:23:37.972502579 +0100 BST
	timestamp: 		 		1405967017
	Milli timestamp: 	 	1405967017000
	Nano timestamp: 	 	1405967017972502579

	$ gotime "2014-02-03 19:54:02"
	-----------------
	UTC: 			 		2014-02-03 19:54:02 +0000 UTC
	Local: 			 		2014-02-03 19:54:02 +0000 GMT
	timestamp: 		 		1391457242
	Milli timestamp: 	 	1391457242000
	Nano timestamp: 	 	1391457242000000000
