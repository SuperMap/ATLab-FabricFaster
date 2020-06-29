
# FabricFaster

Based on Hyperledger Fabric v1.4.4

优化思路是，在一个peer中启动多个链码容器，让交易并行执行，因为我在测试时发现交易流程中，背书阶段的模拟交易是最耗时的。目前调通了链码并行的功能，在我的pc机上测试，使用SDK自己写的测试程序，结果如下：

|           | 查询TPS | 查询提升比例 | 写入TPS | 写入提升比例 |
|-----------|-------|--------|-------|--------|
| 原始（单链码执行） | 86    |        | 35    |        |
| 改进-2链码并行  | 108   | 26%    | 71    | 103%   |
| 改进-4链码并行  | 129   | 50%    | 101   | 189%   |
| 改进-8链码并行  | 168   | 95%    | 118   | 237%   |

## 目前该项目是为了验证思路，不能完全保证正确性和有效性。

除非给我一个理由，否则不知道会不会继续完善。 :sweat_smile: 

## License <a name="license"></a>

Source code files are made available under the Apache License, Version 2.0 (Apache-2.0), located in the LICENSE file. 