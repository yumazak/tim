use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn prints_iana_timezones() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["tz"])
        .assert()
        .success()
        .stdout(predicate::str::contains("Asia/Tokyo"))
        .stdout(predicate::str::contains("UTC"));
}
