# Insurance Company
  - needs to raise money
  - exposed to a risk of a disaster when it needs to pay out claims

- Issues a _catastrophe bond_ 
  - pays attractive rate
  - principal is forgiven if a disaster happens

# Investor
  - needs to invest money
  - needs to diversify with an instrument less dependent on market conditions

How do Issuer and Investor meet their needs?
They use a trusted _third party_: an Investment Bank

# Investment Bank
  - advises on terms
  - connets with investors
  - keeps records
  - forwards payments
  - administers the catastrophe trigger
  - *charges fee*

Out of these tasks what does Investment Bank do that software cannot do?
> tasks except the first one fade out 

- Investment Bank
  - advises on terms

That's what the IB is good for.

# Problems
  - manual process: takes weeks to put together an issue
  - disparate systems: error prone
  - opaque: no open market
  - illiquid: investors cannot trade the bond

# Trust
While software can execute the tasks of a _custodian_ it needs to be trusted by Issuer and Investor at least as much as they trust Investment Bank.

Enter Blockchain
 - shared immutable ledger: the records are kept by all involved parties: Issuers and Investors
 - consensus: the records cannot deviate

# Automation
Blockchain not only keeps records but executes smart contracts that automate the functions of a financial instrument.

- Open market: Issuer and Investor discover each other on the blockchain
- Terms of the bond contract are codified in a smart contract
- Coupon payments are triggered by the smart contract
- Maturity and payment of principal is automated by the contract
- Catastrophe event is automated on a trigger by a public record

# Confidentiality
Hyperledger provides for confidentiality of transactions unlike public blockchains

- Terms of the agreement are visible to counterparties only
- Only vetted parties are allowed to participate

# Cost efficient
Minimum investment in infrastructure or software. 

- Built in security by advanced cryptography. The architecture of blockchain is safer than any single enterprise system or a combination of them.
- Blockchain peer software is deployed at every participant's server
- Low cpu load. Unlike public blockchcains where peers run intensive and wasteful proof-of-work calculations
- No single point of failure

# How does it work in practice?

## Discovery
- Issuer and Investors install blockchain software 
- An authority run either by the Issuer or a regulator enrolls market participants
- Issuer deploys a smart contract with a new bond terms
- Investors query the blockchain, discover the bond
- Investors buy the bond: invoke the bond smart contract with their identity and enter into an agreement with the Issuer

## Payments
The blockchain does not run any cryptocurrency but relies on traditional payment systems to tell it that a money trasnfer occurred. This payment systems are agreed to be trusted by members and can be ACH or SWIFT.

### Bond contract is purchased








