gotime
======

Give hime a timestamp or nano timestamp he will give you the time in a readable format

##Installation

`
go get github.com/juliendsv/gotime
`

##Usage


	$ gotime
	-----------------
	Now: 			 2014-07-22 18:46:30.82667704 +0100 BST
	timestamp: 		 1406051190
	Nano timestamp: 	 1406051190826677040


	$ gotime 1349034753
	-----------------
	Now: 			 2012-09-30 20:52:33 +0100 BST
	timestamp: 		 1349034753
	Nano timestamp: 	 1349034753000000000


	$ gotime 1405967017972502579
	-----------------
	Now: 			 2014-07-21 19:23:37.972502579 +0100 BST
	timestamp: 		 1405967017
	Nano timestamp: 	 1405967017972502579
