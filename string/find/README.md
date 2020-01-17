# 字符串匹配

## KMP
KMP是朴素匹配方法的改进，能够知道在匹配失败后，有多少字符是不需要进行匹配可以直接跳过的，操作方式是逐个移位比较搜索词p和文本s，当搜索词和文本不匹配时，假设p[i] != s[j]，那么p[i-1] == s[j-1]，找到一个索引k，使得p[0:k] == p[i-k:i] (k < i)，这样可以一次移位i-k。

## Rabin-Karp
算法基本思路是计算两个字符串的哈希值，通过比较哈希值的大小来判断是否匹配。

哈希公式如下(base是一个很大的质数):

hash(t[0,m-1]) = t[0] * base<sup>m-1</sup> + t[1] * base<sup>m-2</sup> + ... + t[m-1] * base<sup>0</sup>

举例字符串"abcabc"，假设base为128

hash("abc") = 97 * 128<sup>2</sup> + 98 * 128<sup>1</sup> + 99 * 128<sup>0</sup>

hash("bca") = hash("abc") * 128 + 97 * 128<sup>0</sup> - 97 * 128<sup>3</sup>

hash("cab") = hash("bca") * 128 + 98 * 128<sup>0</sup> - 98 * 128<sup>3</sup>