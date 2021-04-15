#Context Strengthen
根据常见的应用场景暂时分为两种类型的context key：      
####框架级别key，由框架包装，开发时可使用但不能覆盖
#####对研发透明的服务基本信息（ip、host等），只在框架层面一次添加不再读取，可以做一个static map
#####全局的缓存sync.Map
####开发级别key，可以在开发过程中自由添加删除
#####trace信息
#####事务信息
#####当前作用域下的缓存sync.Map