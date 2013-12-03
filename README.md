go-reportman
============

A concurrent application to get financial data from yahoo. This is essentially a CLI to get
csv data for prices. The problem solved here is that R is single threaded and can be pretty
slow when we're trying to obtain large quantities of data. R does shine in doing analysis so
any data gathered from yahoo here can just be used with R.
