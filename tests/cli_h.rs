use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn converts_hour_with_defaults() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h", "9"])
        .assert()
        .success()
        .stdout("0\n");
}

#[test]
fn converts_hour_with_explicit_iana() {
    Command::cargo_bin("tim")
        .unwrap()
        .args([
            "h",
            "--from",
            "America/New_York",
            "--to",
            "Asia/Tokyo",
            "12",
        ])
        .assert()
        .success()
        .stdout(predicate::str::is_match(r"^\d+\n$").unwrap());
}

#[test]
fn converts_hour_with_fractional_offset_tz() {
    // Asia/Kathmandu is +05:45
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h", "--from", "Asia/Kathmandu", "--to", "UTC", "6"])
        .assert()
        .success()
        .stdout("0\n");
}

#[test]
fn converts_hour_with_short_options() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h", "-f", "UTC", "-t", "Asia/Tokyo", "15"])
        .assert()
        .success()
        .stdout("0\n");
}
