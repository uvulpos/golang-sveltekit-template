# Create a new Database Migration

> **SUMMARY:**  
> In this how-to, you'll learn how to create a new user table, how to insert it into the database and what the idea behind agile database migrations (agile devops) is.

## 1. What are agile database migrations?

You have certainly seen a diagram like this before where this circle-ish or snakish kinda diagram shows you an ongoing process, but in really rare cases it is needed to downgrade an application in order to ensure availability.

![agile vs devops graphic](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.ytimg.com%2Fvi%2FTuGdJXfwRBY%2Fmaxresdefault.jpg&f=1&nofb=1&ipt=dabc18f621216755e1136bc801217039fe789f8612aeadc2b33e53555c150a0b&ipo=images)

So you need some sort of downgrade migration to break data relationsships that do not exists yet in our data models and might break the application.

You don't have to use it, but having it is better than needing it

## 2. How to create a new table

Go to your backend service and then to `src/migrator/migration-files` and create two new files: `<number>_<name>.up.sql` and `<number>_<name>.down.sql`. The number and description must match. Just increment the number for the following migrations.

## 3. Execute the migration

If you use the makefiles commands, it gets executed by default. If you execute the binary, use the build in command `app migrate up` to execute the migration scripts
