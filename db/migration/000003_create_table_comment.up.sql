CREATE TABLE public."Comment" (
                                  id BIGSERIAL NOT NULL,
                                  user_id BIGINT NOT NULL,
                                  post_id BIGINT NOT NULL,
                                  "content" varchar NOT NULL,
                                  likes BIGINT DEFAULT 0 NULL,
                                  created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
                                  CONSTRAINT "Comment_pkey" PRIMARY KEY (id)
);


ALTER TABLE public."Comment" ADD CONSTRAINT "Comment_post_id_fkey" FOREIGN KEY (post_id) REFERENCES public."Post"(id) ON DELETE CASCADE;
ALTER TABLE public."Comment" ADD CONSTRAINT "Comment_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."User"(id) ON DELETE CASCADE;