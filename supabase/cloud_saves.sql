-- FNAF Versus V55.0. Run in the SQL editor of the project's Supabase instance.
-- A save is private to its authenticated owner. No public or service-role key in the game.
begin;
create table if not exists public.fnaf_saves (
  user_id uuid primary key references auth.users(id) on delete cascade,
  payload jsonb not null check (jsonb_typeof(payload) = 'object' and payload->>'schema' = '1' and octet_length(payload::text) <= 200000),
  revision integer not null default 1 check (revision > 0),
  updated_at timestamptz not null default now()
);
alter table public.fnaf_saves enable row level security;
revoke all on public.fnaf_saves from anon, authenticated;
grant select, insert, update on public.fnaf_saves to authenticated;
drop policy if exists fnaf_save_select on public.fnaf_saves;
drop policy if exists fnaf_save_insert on public.fnaf_saves;
drop policy if exists fnaf_save_update on public.fnaf_saves;
create policy fnaf_save_select on public.fnaf_saves for select to authenticated using ((select auth.uid()) = user_id);
create policy fnaf_save_insert on public.fnaf_saves for insert to authenticated with check ((select auth.uid()) = user_id);
create policy fnaf_save_update on public.fnaf_saves for update to authenticated using ((select auth.uid()) = user_id) with check ((select auth.uid()) = user_id);
create or replace function public.fnaf_save_revision_guard() returns trigger
language plpgsql set search_path = '' as $$
begin
  if tg_op = 'INSERT' then
    if new.revision <> 1 then raise exception 'invalid initial revision'; end if;
  else
    if new.user_id <> old.user_id or new.revision <> old.revision + 1 then
      raise exception 'save revision conflict';
    end if;
  end if;
  new.updated_at := now();
  return new;
end;
$$;
drop trigger if exists fnaf_save_revision_guard on public.fnaf_saves;
create trigger fnaf_save_revision_guard before insert or update on public.fnaf_saves
for each row execute function public.fnaf_save_revision_guard();
commit;
