# Contributing

Here you'll see some steps and guide lines for contributing to arosh. It's easy and fun. I promisse.

First take a look at the [Architecture](./ARCHITECTURE.md) to get some understanding on how the code base is organized.

This documents should tell all you need to know. If it doesn't, it definely should, so let me know.

# Setting up

Assuming you already have an github account, click some buttons and fork the project. Then clone in your machine:

```bash
git clone https://github.com/YOU/arosh
```

> You might want to use SSH here.

There is some stuff you need to install as well:

- go

That's it. Should be easy to install using your favorite package manager. If you are on `Windows` and outside
the terminal, you are on your own.

After this just make your changes and create a `pull request`.

# Guidelines

- **Commits:** Nothing fancy. Read [Convetional Commits](https://www.conventionalcommits.org/en/v1.0.0/) and you should be good. You may also
  append a scope to the commit: `feat: some` -> `feat(scope): some`. The scope should be the module you'r working on. See [Architecture](./ARCHITECTURE.md).
- **Formatting:** We are using [golines](https://github.com/segmentio/golines) here. Look at the documentation on their repository, install and set your
  favorite editor up. There are no custom configuration for now. The default is good enough.

> **TODO:** A script to format everything 😅.

- **Testing:** If it's new code, then it should be tested. I'm still figuring out how to test some functions on the `lineEditor`, so I won't be mad if you
  struggle as well. The overall approch is: create a array with test cases and iterate through. If it goes wrong print something.
