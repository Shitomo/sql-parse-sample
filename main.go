package main

import (
	"fmt"

	"github.com/zhuxiujia/GoMybatis"
	"github.com/zhuxiujia/GoMybatis/engines"
	"github.com/zhuxiujia/GoMybatis/lib/github.com/beevik/etree"
)

func main() {
	var mapper = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper>
    <!--List<Activity> selectByCondition(@Param("name") String name,@Param("startTime") Date startTime,@Param("endTime") Date endTime,@Param("index") Integer index,@Param("size") Integer size);-->
    <!-- 后台查询产品 -->
    <select id="selectByCondition">
        select * from biz_activity where delete_flag=1
        <if test="name != nil">
            and name like concat('%',#{name},'%')
        </if>
        <if test="startTime != nil">
            and create_time >= #{startTime}
        </if>
        <if test="endTime != nil">
            and create_time &lt;= #{endTime}
        </if>
        order by create_time desc
        <if test="page >= 0 and size != 0">limit #{page}, #{size}</if>
    </select>
</mapper>`

	var builder = GoMybatis.GoMybatisSqlBuilder{}.New(GoMybatis.ExpressionEngineProxy{}.New(&engines.ExpressionEngineGoExpress{}, true), &GoMybatis.LogStandard{}, false)

	var mapperTree = GoMybatis.LoadMapperXml([]byte(mapper))

	fmt.Printf("mapperTree: %v\n", mapperTree)

	var nodes = builder.NodeParser().Parser(mapperTree["selectByCondition"].(*etree.Element).Child)

	fmt.Printf("nodes: %v\n", nodes)

}
