gotime
======

Give it a timestamp or a nano-timestamp it will give you the time in a readable format.

##Installation


	go get github.com/juliendsv/gotime


##Usage


	$ gotime
	-----------------
	UTC: 			 2014-10-08 22:19:09.328685713 +0000 UTC
	Local: 			 2014-10-08 18:19:09.328685713 -0400 EDT
	timestamp: 		 1412806749
	Nano timestamp: 	 1412806749328685713

	$ gotime 1349034753
	-----------------
	UTC: 			 2012-09-30 19:52:33 +0000 UTC
	Local: 			 2012-09-30 15:52:33 -0400 EDT
	timestamp: 		 1349034753
	Nano timestamp: 	 1349034753000000000

	$ gotime 1405967017972502579
	-----------------
	UTC: 			 2014-07-21 18:23:37.972502579 +0000 UTC
	Local: 			 2014-07-21 14:23:37.972502579 -0400 EDT
	timestamp: 		 1405967017
	Nano timestamp: 	 1405967017972502579

	$ gotime "2014-02-03 19:54:02"
	-----------------
	UTC: 			 2014-02-03 19:54:02 +0000 UTC
	Local: 			 2014-02-03 14:54:02 -0500 EST
	timestamp: 		 1391457242
	Nano timestamp: 	 1391457242000000000

	$ gotime yesterday
	-----------------
	UTC: 			 2015-01-16 14:35:34.936662557 +0000 UTC
	Local: 			 2015-01-16 14:35:34.936662557 +0000 GMT
	timestamp: 		 1421418934
	Nano timestamp: 	 1421418934936662557

	$ gotime tomorrow
	-----------------
	UTC: 			 2015-01-18 14:35:34.936748828 +0000 UTC
	Local: 			 2015-01-18 14:35:34.936748828 +0000 GMT
	timestamp: 		 1421591734
	Nano timestamp: 	 1421591734936748828
