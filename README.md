gotime
======

Give it a timestamp or a nano-timestamp it will give you the time in a readable format.

##Installation


	go get github.com/juliendsv/gotime


##Usage


	$ gotime
	-----------------
	UTC: 			 	 2015-01-24 23:50:00.768050008 +0000 UTC
	Local: 			 	 2015-01-24 23:50:00.768050008 +0000 GMT
	timestamp: 		 	 1422143400
	Milli timestamp: 	 1422143400000
	Nano timestamp: 	 1422143400768050008

	$ gotime 1349034753
	-----------------
	UTC: 			 		2012-09-30 19:52:33 +0000 UTC
	Local: 			 		2012-09-30 20:52:33 +0100 BST
	timestamp: 		 		1349034753
	Milli timestamp: 	 	1349034753000
	Nano timestamp: 	 	1349034753000000000

	$ gotime yesterday
	-----------------
	UTC: 			 	 2015-01-23 23:50:00.768304081 +0000 UTC
	Local: 			 	 2015-01-23 23:50:00.768304081 +0000 GMT
	timestamp: 		 	 1422057000
	Milli timestamp: 	 1422057000000
	Nano timestamp: 	 1422057000768304081

	$ gotime -z yesterday
	-----------------
	UTC: 			 	 2015-01-23 00:00:00 +0000 UTC
	Local: 			 	 2015-01-23 00:00:00 +0000 GMT
	timestamp: 		 	 1421971200
	Milli timestamp: 	 1421971200000
	Nano timestamp: 	 1421971200000000000

	$ gotime tomorrow
	-----------------
	UTC: 			     2015-01-25 23:50:00.768348762 +0000 UTC
	Local: 			 	 2015-01-25 23:50:00.768348762 +0000 GMT
	timestamp: 		 	 1422229800
	Milli timestamp: 	 1422229800000
	Nano timestamp: 	 1422229800768348762

	$ gotime 5daysago
	-----------------
	UTC: 			 	 2015-01-19 23:50:00.768468249 +0000 UTC
	Local: 			 	 2015-01-19 23:50:00.768468249 +0000 GMT
	timestamp: 		 	 1421711400
	Milli timestamp: 	 1421711400000
	Nano timestamp: 	 1421711400768468249

	$ gotime 2daysfwd
	-----------------
	UTC: 				 2015-01-29 23:50:00.768497912 +0000 UTC
	Local: 				 2015-01-29 23:50:00.768497912 +0000 GMT
	timestamp: 			 1422575400
	Milli timestamp: 	 1422575400000
	Nano timestamp: 	 1422575400768497912

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
