
-- 記事データを格納するためのテーブル
create table if not exists articles (
  article_id serial primary key,
  title varchar(100) not null,
  contents text not null,
  username varchar(100) not null,
  nice integer not null,
  created_at timestamp
);

-- コメントデータを格納するためのテーブル
create table if not exists comments (
  comment_id serial primary key,
  article_id integer not null,
  message text not null,
  created_at timestamp,
  foreign key (article_id) references articles(article_id)
);