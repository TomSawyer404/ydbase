# What is it?

A simple memory database. It's nothing but a homework to learn primary datastruct of golang.  

# What does it look like?

use `make` to build binary, and then run `./bin/ydbase` to use it.
Default login account is `root 123456`, let's try it:

```bash
# 02-database on 🌱 main [?] via 🐹 v1.17.2 
❯ ./bin/ydbase 
#> root 123456
	Welcome to use ydbase 1.0
$> init 4
Operation success, 4 rows in total
$> select age<=40
Operation success, match 4 rows in total
---------------------------------------------------------------------------
|Order         |Id            |Name          |Age           |Height        
---------------------------------------------------------------------------
|0             |091           |Name1         |23            |1.70          
---------------------------------------------------------------------------
|1             |092           |Name2         |24            |1.70          
---------------------------------------------------------------------------
|2             |093           |Name3         |25            |1.70          
---------------------------------------------------------------------------
|3             |094           |Name4         |26            |1.70          
---------------------------------------------------------------------------
Total print 4 row(s)
---------------------------------------------------------------------------
$> count
Operation success, 4 rows in total
$> delete 091
Operation success, 3 rows in total
$> insert 001,jack,33,1.81
Operation success, 4 rows in total
$> select name==jack
Operation success, match 1 rows in total
---------------------------------------------------------------------------
|Order         |Id            |Name          |Age           |Height        
---------------------------------------------------------------------------
|0             |001           |jack          |33            |1.81          
---------------------------------------------------------------------------
Total print 1 row(s)
---------------------------------------------------------------------------
$> update 001,Alice,18,1.79
Operation success, 1 row affected
$> select name%A
Operation success, match 1 rows in total
---------------------------------------------------------------------------
|Order         |Id            |Name          |Age           |Height        
---------------------------------------------------------------------------
|0             |001           |Alice         |18            |1.79          
---------------------------------------------------------------------------
Total print 1 row(s)
---------------------------------------------------------------------------
$> exit
See you!

```

# References

The homework is inspired by Chapter 5 of [this book](http://product.dangdang.com/29120162.html)

---
