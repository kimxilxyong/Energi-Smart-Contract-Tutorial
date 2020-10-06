# Writing smartcontract DApps on the Energi3 blockchain

## Part X : Installing Svelte with Framework7 gui framework

### 1. Install the Framework7

Go to or make directory svelte-framework7-ListInputPassword and run theese commands:

```bash
> npm install framework7 framework7-cli
...
> npm install -g npx
...
> npx framework7-cli create

Select: PWA
Select: framework svelte
Select: Single View
```

-

### 1. Use the App generator started in the last step

Framework7 is a wonderfull library for svelte, which we will use for the frontend client implementation.

Select PWA and then svelte as framework (NOT react, NOT vue!!)

You should now have a starter App in svelte-framework7-ListInputPassword, build it:
Webpack is used as a bundler.

Compile:

```bash
npm run build-dev
```

Run:

```bash
npm run dev
```

-

### 3. Install the Svelte extensions for VS Code

Pic Extensions
