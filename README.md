check-receiver
==============

A small utility which receives calls via HTTP-Post when a tasks on some host succeeds (ie a backup).
It then touches a file corresbonding to the path POSTed to.
This file is then checked with a monit-check.

This way monit can check asyncronus tasks that run, ie. every day.
