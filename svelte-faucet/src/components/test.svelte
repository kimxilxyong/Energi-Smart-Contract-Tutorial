<!-- SPDX-License-Identifier: MIT
     Copyright (c) 2020 Kim Il Yong -->

<Page name="web3" onPageBeforeIn="{onShow}" onPageBeforeOut="{onHide}">
            <ListInput
              outline
              label="Web3 Provider"
              floatingLabel
              type="text"
              clearButton
              value="{web3URL}"
              onInputClear="{() => {
                web3URL = '';
              }}"
              onBlur="{(e) => urlChanged(e)}"
              style="vertical-align:top;margin-top:1vh;"
              spellcheck="false"
            >
              <!-- <i class="icon demo-list-icon" slot="media" />   margin-top:24px -->
            </ListInput>
          <Button
            fill
            raised
            on:click="{ConnectToRPCProviderUrl}"
            style="vertical-align:top;margin-top:2.8vh;bgcolor:var(--energi-color-green)"
          >
            <i class="f7-icons size-14" style="margin-right: 5%">link</i>
            Connect
          </Button>
                    <ListInputPassword
                      label="Keystore File Password"
                      clearButton="{true}"
                      bind:password
                    />


</Page>

<script>
  import Web3 from "web3";
  import {
    Page,
    ListInput,
    Button,
    f7
  } from "framework7-svelte";

  import ListInputPassword from "../components/input-password.svelte";


</script>

