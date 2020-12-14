<!-- SPDX-License-Identifier: MIT
     Copyright (c) 2020 Kim Il Yong
     Version 1.0.0 -->

<ListInput
    input={true}
    outline={outline}
    label="{label}"
    floatingLabel="{floatingLabel}"
    type="{passwordType}"
    clearButton="{clearButton}"
    id="{id}"
    inputId="REMOVEDINPUT"
    value="{password}"
    onInputClear="{() => {password = ""; clearButtonVisible = false; thisInput.focus();}}"
    onFocus={onFocus}
    onBlur={onBlur}
    onInputNotempty={onInputNotempty}
    onInputEmpty={onInputEmpty}
  >
      <div slot="input" style="display:flex;flex-direction: row;justify-content: flex-end;align-items:stretch;min-height: inherit;">
          <div style="flex: 1 1 90%;">
            {#if passwordVisible}
                <input type="text" bind:value="{password}" bind:this={thisInput} id={inputId}
                on:input="{() => onInput()}"
                on:change="{() => onChange()}"
                spellcheck="false"
                >
            {:else}
                <input type="password" bind:value="{password}" bind:this={thisInput} id={inputId}
                on:input="{() => onInput()}"
                on:change="{() => onChange()}"
                spellcheck="false"
                >
            {/if}
          </div>
          <div style="max-width:2.5em;min-width:2em;margin-right:6px;flex: 0 0 5%;display:flex;justify-content:center;align-items:center;z-index:12;cursor:pointer;">
            <div class="f7-icons password-eye-button" on:click="{() => togglePasswordHidden()}">{eyeStyle}</div>
          </div>
          {#if clearButtonVisible}
            <div style="max-width:1.7em;min-width:1.6em;flex: 0 0 5%;justify-content:center;align-items:center;">
            </div>
          {/if}
      </div>
</ListInput>

<style>
    .password-eye-button
    {
        z-index: 11;
        color: grey;
        font-size: 22px;
        transition-duration: 100ms;
        pointer-events: visible;
        -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
        -webkit-font-smoothing: antialiased;
    }
</style>

<script>
    import {
      ListInput,
    } from 'framework7-svelte';
    import { tick } from 'svelte';

    export let passwordVisible = false;
    export let password;
    export let label;
    export let floatingLabel = true;
    export let clearButton = false;
    export let outline = true;
    export let id;
    export let inputId;
    export let clearButtonVisible = false;
    export let onFocus;
    export let onBlur;
    export let onInputNotempty;
    export let onInputEmpty;

    let thisInput;
    let passwordType = "password";
    let eyeStyle = "eye";

    /* eslint-ignore next line */
    async function togglePasswordHidden() {
      if (passwordVisible) {
        passwordType = "password";
        passwordVisible = false;
        eyeStyle = "eye";
      } else {
        passwordType = "text";
        passwordVisible = true;
        eyeStyle = "eye_slash";
      }
      await tick();
      thisInput.focus();
    }
    function onInput() {
      thisInput.blur();
      thisInput.focus();
    }
    function onChange() {
      if ( thisInput.className.includes("input-with-value")) {
        clearButtonVisible = true;
      } else {
        clearButtonVisible = false;
      }
    }
  </script>
