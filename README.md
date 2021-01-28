#### LRU实现及测试

LRU实现有下面三种方式： 
1. Hash+双链表的实现
2. Hash+数组实现
3. Hash+双链表+sync.Mutex实现的并发LRU

相关的测试数据:

```
=== RUN   TestConcurrentSafeLRU
--- PASS: TestConcurrentSafeLRU (0.00s)
=== RUN   TestNewDoubleLinkList
--- PASS: TestNewDoubleLinkList (0.00s)
=== RUN   TestDoubleLinkList_Remove
--- PASS: TestDoubleLinkList_Remove (0.00s)
=== RUN   TestLinkListLRU
--- PASS: TestLinkListLRU (0.00s)
=== RUN   TestSliceLRU
--- PASS: TestSliceLRU (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/zhangmingkai4315/liblru
BenchmarkLRU_Set/slicelru-16-6         	 5643759	       212 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-32-6         	 5608861	       214 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-64-6         	 5720314	       216 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-128-6        	 5598890	       213 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-256-6        	 5601728	       214 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-512-6        	 5552394	       216 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-1024-6       	 5451608	       221 ns/op	      81 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-2048-6       	 5244849	       232 ns/op	     105 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-4096-6       	 4993042	       241 ns/op	     105 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-8192-6       	 4715394	       258 ns/op	     105 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-16384-6      	 4395130	       265 ns/op	     104 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-32768-6      	 4266280	       271 ns/op	     104 B/op	       3 allocs/op
BenchmarkLRU_Set/slicelru-65536-6      	 4034008	       304 ns/op	     104 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-16-6      	 5975928	       202 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-32-6      	 5873912	       205 ns/op	      66 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-64-6      	 5806636	       207 ns/op	      66 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-128-6     	 5822494	       209 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-256-6     	 5749809	       211 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-512-6     	 5665566	       214 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-1024-6    	 5512400	       218 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-2048-6    	 5381076	       223 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-4096-6    	 5176992	       234 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-8192-6    	 4872079	       248 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-16384-6   	 4757466	       258 ns/op	      64 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-32768-6   	 4722590	       260 ns/op	      64 B/op	       3 allocs/op
BenchmarkLRU_Set/linklistlru-65536-6   	 4555132	       278 ns/op	      64 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-16-6    	 4915743	       246 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-32-6    	 4965199	       247 ns/op	      66 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-64-6    	 4971882	       246 ns/op	      66 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-128-6   	 4876923	       247 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-256-6   	 4803801	       253 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-512-6   	 4732929	       254 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-1024-6  	 4698464	       256 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-2048-6  	 4530540	       265 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-4096-6  	 4432050	       271 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-8192-6  	 4227070	       287 ns/op	      65 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-16384-6 	 4058554	       297 ns/op	      64 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-32768-6 	 4075593	       295 ns/op	      64 B/op	       3 allocs/op
BenchmarkLRU_Set/concurrentlru-65536-6 	 3998179	       312 ns/op	      64 B/op	       3 allocs/op
BenchmarkLRU_GetInLRU/slicelru-16-6    	32726574	        36.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-32-6    	30145534	        40.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-64-6    	27706000	        44.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-128-6   	20857371	        58.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-256-6   	13678526	        82.2 ns/op	       1 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-512-6   	10454690	       110 ns/op	       2 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-1024-6  	 8240058	       145 ns/op	       2 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-2048-6  	 5471119	       219 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-4096-6  	 2790309	       423 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-8192-6  	 1259122	       934 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-16384-6 	  612566	      1919 ns/op	       4 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-32768-6 	  274688	      4876 ns/op	       4 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/slicelru-65536-6 	   92492	     18678 ns/op	       5 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-16-6 	44946888	        26.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-32-6 	44497917	        27.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-64-6 	43226445	        27.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-128-6         	33313114	        36.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-256-6         	22735588	        53.4 ns/op	       1 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-512-6         	17907459	        66.3 ns/op	       2 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-1024-6        	18012878	        67.3 ns/op	       2 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-2048-6        	17022465	        70.3 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-4096-6        	16327630	        73.6 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-8192-6        	15526083	        78.4 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-16384-6       	14638765	        80.6 ns/op	       4 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-32768-6       	14201324	        84.2 ns/op	       4 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/linklistlru-65536-6       	13548398	        88.5 ns/op	       5 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-16-6        	18037994	        65.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-32-6        	18812254	        64.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-64-6        	18194116	        65.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-128-6       	16386531	        73.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-256-6       	12958563	        92.1 ns/op	       1 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-512-6       	11837366	       100 ns/op	       2 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-1024-6      	11567442	       103 ns/op	       2 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-2048-6      	11228350	       107 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-4096-6      	10915854	       110 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-8192-6      	10640328	       114 ns/op	       3 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-16384-6     	10199158	       118 ns/op	       4 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-32768-6     	 9861234	       120 ns/op	       4 B/op	       0 allocs/op
BenchmarkLRU_GetInLRU/concurrentlru-65536-6     	 9811418	       123 ns/op	       5 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-16-6          	46393736	        25.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-32-6          	37216833	        32.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-64-6          	32763439	        36.9 ns/op	       1 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-128-6         	19697925	        61.1 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-256-6         	17552253	        68.9 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-512-6         	17182101	        69.8 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-1024-6        	23831101	        50.9 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-2048-6        	16556961	        72.9 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-4096-6        	15297328	        77.1 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-8192-6        	19120440	        61.9 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-16384-6       	14090232	        83.8 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-32768-6       	13062770	        88.9 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/slicelru-65536-6       	15949603	        75.7 ns/op	       6 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-16-6       	49336882	        24.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-32-6       	50373553	        24.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-64-6       	34269847	        35.0 ns/op	       1 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-128-6      	19938996	        60.4 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-256-6      	17381228	        68.9 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-512-6      	17403937	        69.3 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-1024-6     	24229489	        50.1 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-2048-6     	16636716	        73.3 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-4096-6     	15588892	        77.7 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-8192-6     	18984390	        62.4 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-16384-6    	14407851	        82.2 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-32768-6    	13212386	        88.3 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/linklistlru-65536-6    	15610710	        75.5 ns/op	       6 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-16-6     	18227547	        64.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-32-6     	17482461	        69.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-64-6     	16065777	        74.7 ns/op	       1 B/op	       0 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-128-6    	12315284	        98.0 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-256-6    	11221873	       107 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-512-6    	11136913	       106 ns/op	       3 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-1024-6   	13306876	        91.6 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-2048-6   	10737057	       113 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-4096-6   	10340419	       116 ns/op	       4 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-8192-6   	12197889	        98.9 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-16384-6  	 9904753	       120 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-32768-6  	 9397945	       128 ns/op	       5 B/op	       1 allocs/op
BenchmarkLRU_GetNotInLRU/concurrentlru-65536-6  	10606626	       113 ns/op	       6 B/op	       1 allocs/op
PASS
ok  	github.com/zhangmingkai4315/liblru	159.588s


```