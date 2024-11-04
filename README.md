![](https://github.com/Gateway-DAO/network-ui/blob/develop/public/social.png)

# Gateway Javascript SDK

[![report-card][report-card]][report-card]
[![Join Discord][discord-image]][discord-url]
[![Run Eslint & Test cases](https://github.com/Gateway-DAO/gateway-go-sdk/actions/workflows/test.yaml/badge.svg)](https://github.com/Gateway-DAO/gateway-go-sdk/actions/workflows/test.yaml)
[![Coverage Status][codecov-image]][codecov-url]

## Introduction

A GO SDK for the Gateway API. We are still developing and testing go sdk. Right now the go sdk is in beta stage and code might change in future. 


### Features

- Full type information for methods and responses.
- Bearer Token Support
- Supports Go 1.22.8+.

## Installation

### Using Go Get

```
go get github.com/Gateway-DAO/gateway-go-sdk/client@v0.0.12-beta
```


## Gateway Client

To setup the gateway client we will authenticate with a bearer-token,or use wallet private key as follows

<!-- ```typescript
import { Gateway, WalletTypeEnum } from '@gateway-dao/sdk';

// using jwt
// you need to manage jwt expration on your own. JWT expires after 5 days once issued
const gateway = new Gateway({
  jwt: 'your-jwt', // store in env file!
});

// using wallet private key
// here we manage jwt on your behalf thus providing better experience
const gateway = new Gateway({
  wallet: {
    privateKey: 'your-private-key', // store in env file!
    walletType: WalletTypeEnum.Ethereum,
  },
});
``` -->

**The wallet private key is not send anywhere and is just used to sign messages on behalf of developer using it to aviod to manage jwt token himself. This way we minimize jwt expiration errors on API and provide smoother developer experience**

**If you are using bearer token make sure you add token without Bearer as we add Bearer automatically when you make request. Else it will give you Unauthorized error even if your token is correct**
For example

<!-- ```typescript
import { Gateway } from '@gateway-dao/sdk';

const gateway = new Gateway({
  token: 'Bearer your-token', // this is wrong won't work
});
``` -->

This library supports Bearer Token along. Do not share your authentication token with people you don’t trust. This gives the user control over your account and they will be able to manage Data Assets (and more) with it. Use environment variables to keep it safe.

## Examples

Make sure to add try catch blocks around methods to catch all the validation and API based errors.

<!-- ### Creating a PDA

```typescript
import { Gateway, AccessLevel, WalletTypeEnum } from '@gateway-dao/sdk';

const gateway = new Gateway({
  wallet: {
    privateKey: 'your-private-key', // store in env file!
    walletType: WalletTypeEnum.Ethereum, // supported types are ethereum,solana,sui
  },
}); -->

<!-- async function main() {
  try {
    const claim = {
      firstName: 'Test',
      age: 21,
    };
    const id = await gateway.dataAsset.createStructured({
      name: 'THIS IS A TEST',
      data_model_id: 508557480951911,
      tags: ['football', 'sports'],
      claim: claim,
      acl: [
        {
          address: 'another-user-did',
          roles: [AccessLevel.VIEW, AccessLevel.SHARE], // assing roles 
        },
      ],
    });
    console.log('\nData Asset created with ID:', id);
  } catch (error) {
    console.log(error); // Can log it for degugging
  }
}

main();
``` -->

<!-- ### Creating a Data Model

```typescript
import { Gateway, AccessLevel, WalletTypeEnum } from '@gateway-dao/sdk';

const gateway = new Gateway({
  wallet: {
    privateKey: 'your-private-key', // store in env file!
    walletType: WalletTypeEnum.Ethereum, // supported types are ethereum,solana,sui
  },
}); -->

<!-- async function main() {
  try {
    const dataModelBody = {
    title: 'test sdk dm',
    description: 'test sdk dm',
    schema: {
      type: 'object',
      title: 'Gateway ID',
      default: {},
      required: [],
      properties: {
        name: {
          type: 'string',
          title: 'Name',
        },
      },
      additionalProperties: false,
    },
  };
  result = await gateway.dataModel.create(dataModelBody);
  } catch (error) {
    console.log(error); // Can log it for degugging
  }
}
main();
``` -->

## More examples

We have created a separate repository which have more [examples you can access it here](https://github.com/Gateway-DAO/sdk-scripts-example/)

## Error Handling

Incase of any API errors we will throw a custom message which is a string which has all neccessary info regarding error. Make sure to use try catch blocks to handle those.

## License

The Gateway GO SDK is licensed under the [MIT License](LICENSE).

## Contributing

If you want to support the active development of the SDK. [Please go through our Contribution guide](docs/CONTRIBUTING.md)

## Code of Conduct

Please read our [Code of Conduct](docs/CODE_OF_CONDUCT.md) before contributing or engaging in discussions.

## Security

If you discover a security vulnerability within this package, please open a ticket on Discord. All security vulnerabilities will be promptly addressed.

## Support

We are always here to help you. Please talk to us on [Discord](https://discord.gg/tgt3KjcHGs) for any queries.

[report-card]: https://goreportcard.com/badge/github.com/Gateway-DAO/gateway-go-sdk
[codecov-image]: https://codecov.io/gh/Gateway-DAO/gateway-go-sdk/graph/badge.svg?token=8N92RFGZHI
[codecov-url]: https://codecov.io/gh/Gateway-DAO/gateway-go-sdk
[discord-image]: https://img.shields.io/discord/733027681184251937.svg?style=flat&label=Join%20Community&color=7289DA
[discord-url]: https://discord.gg/tgt3KjcHGs