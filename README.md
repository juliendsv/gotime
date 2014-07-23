gotime
======

Give it a timestamp or a nano-timestamp it will give you the time in a readable format.

##Installation


	go get github.com/juliendsv/gotime


##Usage


	$ gotime
	-----------------
	Date: 			 2014-07-23 18:36:08.319846333 +0100 BST
	timestamp: 		 1406136968
	Nano timestamp: 	 1406136968319846333


	$ gotime 1349034753
	-----------------
	Date: 			 2012-09-30 20:52:33 +0100 BST
	timestamp: 		 1349034753
	Nano timestamp: 	 1349034753000000000


	$ gotime 1405967017972502579
	-----------------
	Date: 			 2014-07-21 19:23:37.972502579 +0100 BST
	timestamp: 		 1405967017
	Nano timestamp: 	 1405967017972502579


	$ gotime "2014-02-03 19:54:02"
	-----------------
	Date: 			 2014-02-03 19:54:02 +0000 GMT
	timestamp: 		 1391457242
	Nano timestamp: 	 1391457242000000000
