// Package exchange would help with parsing the provided by National Bank rates.
// Feed at:
// https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml
package exchange

/*
TODO:
- add proper logging including log levels, example of library I wrote:
https://github.com/TudorHulban/log

- add validation that day is not older than 90 days

- add decrement for day in order to pick up the rate for when today is
holiday and we do not have a oficial rate

- add caching. maybe we should not go at every request and fetch but keep data locally
*/
