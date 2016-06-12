# Catastrophe Bond Proof Of Concept
This is a functional specification document for the Proof Of Concept of a system enabling insurance companies to issue Catastrophe Bonds on Hyperledger blockchain.

Best viewed with [Markdown Reader Plugin for Chrome](https://chrome.google.com/webstore/detail/markdown-reader/gpoigdifkoadgajcincpilkjmejcaanc?utm_source=chrome-app-launcher-info-dialog).  
Present with [Markdown Slides Plugin for Chrome](https://chrome.google.com/webstore/detail/markdown-slides/ndpkdbdonkcnidnnmdgphflfcojnhoaa?utm_source=chrome-app-launcher-info-dialog).

---
# Catastrophe Bond

> Catastrophe bonds emerged from a need by insurance companies to alleviate some of the risks they would face if a major catastrophe occurred, which would incur damages that they could not cover by the premiums, and returns from investments using the premiums, that they received.  
An insurance company issues bonds through an investment bank, which are then sold to investors. These bonds are inherently risky, generally BB, and usually have maturities less than 3 years. If no catastrophe occurred, the insurance company would pay a coupon to the investors, who made a healthy return. On the contrary, if a catastrophe did occur, then the principal would be forgiven and the insurance company would use this money to pay their claim-holders.  
Investors include hedge funds, catastrophe-oriented funds, and asset managers. They are often structured as floating-rate bonds whose principal is lost if specified trigger conditions are met. If triggered the principal is paid to the sponsor. The triggers are linked to major natural catastrophes. Catastrophe bonds are typically used by insurers as an alternative to traditional catastrophe reinsurance.

https://en.wikipedia.org/wiki/Catastrophe_bond

---
# Problem

* An insurance company contracts an intermediary such as an investment bank to issue bonds on their behalf and act as a custodian of bond contracts.

* Bonds are bought directly from the insurance company by _subscribers_: investors who hold bond contracts to earn coupon payments. The insurance company would like to allow investors to trade contracts to attract more investors on better terms.

---
# Solution

* Permissioned Blockchain
  * immutable ledger
  * no third party and not a centralized solution
  * run and trusted by all members
  * identity obfuscated to other members
  * transparent to auditors

* Chaincode
  * software executed on the blockchain trusted to change records on the ledger
  * members execute chaincode functions based on their roles
  * ledger records are updated based on the identity of members

---
# Records

Records are stored on the blockchain either as the chaincode's key value sets or tables.

* Bond  
  a record in the chaincode's state of a bond issued by one of the members. Encapsulates properties common for bond issue such as term, rate and catastrophe trigger.

* Contract  
  bonds are divided into contracts to sell to investors and trade. All contracts share attributes of their bond such as term and coupon. Contracts are digital assets that cannot be divided and are owned either by their issuer or investor. When investors trade, the chaincode transfers ownership from seller to buyer.  
  Represents a state machine where it is 
  * created by the chaincode when a bond is issued
  * then bought from issuers by investors
  * then traded among investors
  * finally canceled when the maturity date of its bond is reached or the catastrophe event is triggered

---

* Trade  
  a record of a change of ownership of a contract. Represents a state machine where 
  * initially it is an `offer` by the current owner to sell a contract at a specified price
  * then a trade is `captured` when a buyer agrees to trade
  * finally the trade is `settled` when the transfer of money compensating the seller triggers change of ownership of the contract

---
# Actors

Actors are users of various roles in the system. Their ids and roles are defined in the blockchain's member services records. Other attributes may also be defined in the member services or on the blockchain or in a separate persistence store. For the POC it's sufficient to predefine a set of users with all of their attributes in the member services `yaml` files.

* Member  
  actors transacting on the blockchain
  * Issuer  
  an insurance company allowed to issue bonds. Issuers receive principal amount when they sell bond contracts to investors and pay the investors monthly coupons.
  * Investor  
  a hedge fund allowed to buy and sell bond contracts to other members. Investors are compensated by issuers by payments of monthly coupons.

* Auditor  
  a regulator allowed to query the blockchain with read only access; does not transact.

---
# Oracles

*Oracle*  is a mechanism to inject events occurring outside of the blockchain. It is a third party agreed to be trusted by all members and operates as a node on the blockchain with special permissions to the chaincode methods.

* Payment Oracle  
  notifies the blockchain that a transfer of fiat money occurred between members' bank accounts. Can be a client listening on events in ACH network.  
  _example_ an ACH client permissioned by the issuer to listen on events in his accounts payable.  
  Calls the chaincode method to notify the blockchain that a payment of a coupon occurred. The chaincode increments the number of coupons paid on the contract.   
  _example_ an ACH client permissioned by the investor to listen on events in his trading account.  
  Notifies the blockchain that a payment for a bond contract sale occurred. The chaincode transfers ownership of the contract to the buyer.  

---

* Catastrophe Oracle  
  notifies the blockchain that a catastrophe condition is triggered  
  _example_ a national weather service reports a hurricane of category 2. The catastrophe oracle calls the chaincode which in turn marks affected bonds `triggered` which stops payment of their coupon and prevents paying back their principal.

---
# Member

All members: issuers and investors have these attributes.

- id  
  unique id of the member organization: issuer or investor. May be chosen by the member.  
  _example_ `issuer01`, `Allianz`, `CatHedgeFund`
- password  
  a secret needs to be passed to the blockchain to authenticate the member to transact
- paymentAccount  
  various payment oracles will identify members by accounts in respective payment systems
  - aba.account  
    composite unique identifier of the member's fiat money bank account used by a bank transfer ACH oracle. When a payment of fiat money occurs the payment oracle will report to the blockchain a money transfer event with parties' bank aba and account numbers.  
    aba  
    member bank identifier   
    account  
    member bank account identifier

---
# Bond

- id  
  composite unique identifier: `issuer member id.year.month.rate in basis points`  
  _example_ `issuer01.21.06.600` identifies a bond issued by issuer01 maturing in June 2021 with 6% annual coupon rate
- ownerId  
  member id of the issuer
- principal  
  total amount borrowed. Usually divided into a number of contracts traded individually.  
  _example_ an issue of a bond with $1,000,000 principal creates 10 contracts of $100,000
- term  
  number of months before the principal must be paid back  
  _example_ a bond issued in June 2016 with a 60 month term will _mature_ in June 2021: the issuer will need to pay back each holder of a contract the principal of $100,000 

---

- rate  
  percentage of the principal paid to investors by the issuer each year  
  _example_ a bond with a rate of 6% and principal of $1,000,000 will pay an investor holding one contract of $100,000 a monthly coupon of `$100,000 * 0.06 / 12 = $500`
- catastrophe  
  a trigger when the principal won't have to be repaid  
  _example_ if a hurricane of category 2 hits Florida during the term of the bond the issuer won't have to pay back the principal to the investors holding bond contracts
- state
  - `active` before maturity date
  - `matured` after maturity date
  - `triggered` when a catastrophe occurred before maturity date

---
# Contract

- id  
  composite unique identifier: `issuer member id.year.month.rate in basis points.contract id`  
  _example_ `issuer01.21.06.600.3` identifies fourth contract for a bond issued by issuer01 maturing in June 2021 with 6% annual coupon  
  contract id  
  sequential identifier unique within an issue  
  _example_ an issue of a bond with $1,000,000 principal creates 10 contracts of $100,000 with ids from 0 to 9
- ownerId  
  member id of the current owner of the contract: an investor holding the contract and collecting coupon payments
- couponsPaid  
  number of coupons already paid to investors affects the contract's price  
  _example_ a bond issued Jan 2016 has 5 coupons paid out to investors by the end of Jun 2016

---
# Trade

- id  
  sequential id unique per chaincode
- contractId  
  id of the contract offered for sale and traded  
- price  
  amount of money that needs to be transferred from buyer to seller in order for the contract to be moved to the buyer  
  expressed as percentage of the contract's face value  
  if the contract is offered for sale the price is set by the current owner: either the issuer or an investor  
  _example_  when the bond is issued its price is set to 100 by the issuer and offered to _subscribers_: the initial investors who will buy the contracts at 100% of their face value  
  _example_ an investor offers a contract for sale of a 60 month 6% bond that has paid 12 coupons already at a price of `$100,000 - $500 * 12 / $100,000 = 94`  
  _example_ an investor holds a contract of a bond whose catastrophe trigger is likely to occur. He offers the contract for sale at the price of 5 reflecting his view of the high probability of the catastrophe occurring. The investor may get $5000 for the $100,000 contract if he manages to sell it or won't get anything if the catastrophe happens and the whole principal is lost.

---
- state
  - `offer` when created
  - `captured` when a buyer agrees to trade
  - `settled` after the transfer of money compensating the seller

---
# Web Application

The POC is presented as a web application with pages offering functionality to 
* Issuers: issue and track bond contracts
* Investors: buy contracts from issuers and trade among themselves
* Auditors: query bonds, contracts and trades

The following sections describe these pages per actor.

---
# Issuer: Inventory

List of contracts the issuer has issued, their owners and the bond's current state.

bond | contract |  owner | coupons paid | state
---|---|---|---|---|
issuer01.21.06.600|1|investor01|3|triggered
issuer01.25.01.550|1|investor01|5|active
issuer01.25.01.550|2|investor02|4|active
issuer01.16.01.650|1|investor03|60|matured
issuer01.16.01.650|2|investor03|60|matured

---
# Issuer: Issue a Bond

When clicking on a plus button on issuer's Inventory page a modal window opens prompting the issuer to enter parameters of a new bond.

ISSUE | Bond |
---|---
principal | $1,000,000
term | 60 months
rate | 6.25%
trigger | (dropdown of `hurricane category 2 Florida` or `earthquake 5 California`)

When details are entered and _Issue_ button is clicked a call is made to the chaincode which creates a record of the new bond and records of its contracts. The contracts are immediately offered for sale and appear on the Market page.

---
# Investor: Inventory

List of contracts an investor currently owns.

bond | contract |  issuer | coupons paid | state
---|---|---|---|---|
issuer01.21.06.600|1|issuer01|3|triggered
issuer01.25.01.550|1|issuer01|5|active

---
# Investor: Offer a Contract For Sale

When clicking on a contract on Investor's Inventory page a modal window comes up prompting the investor to enter price and offer the contract for sale. When _Sell_ button is clicked a call is made to the chaincode's method that creates a trade record in _offer_ state. The Market page shows this new offer.

SELL | Contract |
---|---
bond | issuer01.21.06.600
contract id | 3
price | 99

---
# Investor: Market

A blotter listing all sale offers by investors and issuers. Investors can choose a contract to buy and initiate a trade.


bond | contract id |  owner | coupons paid | price
---|---|---|---|---|
issuer01.21.06.600|3|issuer01|3|99
issuer01.25.01.550|4|issuer01|5|97
issuer01.25.01.550|2|investor02|4|98
issuer01.16.01.650|1|investor03|60|90
issuer01.16.01.650|2|investor03|60|91

---
# Investor: Buy Contract

When clicking on a sale offer on Market page a modal window opens prompting the investor to agree to a trade.

BUY | Contract |
---|---
from | issuer01
bond | issuer01.21.06.600
contract id | 3
price | 99

When _Buy_ button is clicked a call is made to the chaincode which updates a record of the trade turning it from _offer_ to _captured_. The trade offer leaves the Market page.

---
# Payment Oracle

For the POC this will be a simple demo page imitating a call from an ACH client informing the blockchain of a money transfer either from a coupon paid to investor or from a buyer to seller of a contract.

TRANSFER | Money |
---|---
from aba.account | 02100210.12345678
to aba.account | 03200320.09876543
amount | $500
purpose | (dropdown of trade or coupon or principal)
id | id of trade or contract

Once the _Transfer_ button is clicked the chaincode's method is invoked with the transfer details entered. The chaincode in turn finds the parties of the transfer and the contract and marks the coupon paid or trade settled.

---
# Catastrophe Oracle

For the POC this will be a simple demo page imitating a call from a weather service client informing the blockchain that a catastrophe occurred.

TRIGGER | Catastrophe |
---|---
catastrophe | (dropdown of `hurricane category 2 Florida` or `earthquake 5 California`)

Once the _Trigger_ button is clicked the chaincode's method is invoked to search for all bonds whose trigger matches the one selected, their contracts are then moved to `triggered` state.

---
# Timing Mechanism

A mechanism: either an oracle or a service offered by the blockchain is to inform the chaincode of progression of time. The chaincode needs to know whether a month has passed and a coupon needs to be paid, or a maturity date has been reached and the bond needs to be marked as matured and coupon payments stopped.

For the POC that service can be imitated by an interval service of a browser calling the chaincode's method to progress time and at the same time displaying current month and year on every page. The progression of time in the demo needs to be accelerated so a month is changed every minute, and controlled via _start_ and _stop_ buttons.

















