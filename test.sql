select id from image_tables;



-- 创建索引
create index index_length on image_tables (id,length);

create index index_data on image_tables (id,filename);

