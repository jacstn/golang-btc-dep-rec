Bitcoin Incoming Transactions Reconciliation

# Installation

## Using Docker
- run command `docker-compose up`. If you face any issue with installation check *Known issues section*.

### Tests
Tests are only available under local installation
  - run `go test` to perform unit testing of application

## Application and Database
Application has been build in `go1.19`
MYSQL database engine has been used in version `8.0.28`

### Database Structure

#### 2 tables were designed to store all informatin required for this code challenge.
  - `Customer` - store personal information about customer and its Bitcoin Address
  - `BTCDeposit` - store information about deposits

## Unique Deposit
Deposits are stored in `BTCDeposit` table and it's uniqueness is protected by composite primary key that includes:
  - transaction ID - `txid`
  - vout - `vout`

## Valid Deposit Assumptions
  - deposit `0.00000000` is *not* a valid deposit
  - deposits has to have minimum 6 confirmations
  - coinbase transactions are treated as deposits, and only valid with 100+ confirmations (`generate` category)
  - when a single transaction has *multiple* deposits on the same address they are treated as *seperated* deposits. If this is not correct, proposed solution would be to merge transactions under new composite key: (`txid`, `address`) and remove `vout` as it would not be necessary.

## Services
  - `data-service` - responsible for deliverying transactions data
  - `btc-transaction-service` - responsible for analyzing and filtering transactions
  - `btc-deposit-info-service` - responsible for analyzing deposits in database and producing report data


# Known Issues
  - Docker installation and `.env`. During docker installation and launching `mysql database` service, database is populated using `database-docker.sql` file. This file has hardcoded database name and in case of changing database name `MYSQL_DB_PROD` in `.env` file `database-docker.sql` must be also updated.

# Room for improvement
 - Cover Application with more tests
