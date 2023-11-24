insert into articles (
  title,
  contents,
  username,
  nice,
  created_at
) values (
  'initial article',
  'thi is initial article!!',
  'nanashino',
  14,
  now()
);

insert into articles (
  title,
  contents,
  username,
  nice,
  created_at
) values (
  '222222 article',
  'this is second article !?',
  'nanashino',
  2323,
  now()
);

insert into comments (
  article_id,
  message,
  created_at
) values (
  1,
  'first comment',
  now()
);

insert into comments (
  article_id,
  message
) values (
  1,
  'first comment'
);