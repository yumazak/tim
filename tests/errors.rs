use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn fails_on_invalid_hour() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h", "25"])
        .assert()
        .failure()
        .stderr(predicate::str::contains("invalid hour"));
}

#[test]
fn fails_on_invalid_timezone() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h", "--from", "Invalid/TZ", "9"])
        .assert()
        .failure()
        .stderr(predicate::str::contains("invalid timezone"));
}

#[test]
fn reports_stdin_errors_to_stderr() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h"])
        .write_stdin("9\nabc\n12\n")
        .assert()
        .code(1)
        .stdout("0\n3\n")
        .stderr(predicate::str::contains("invalid hour"));
}

#[test]
fn fails_on_nonexistent_dst_datetime() {
    // 2024-03-10 02:30 doesn't exist in America/New_York (spring forward)
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt", "--from", "America/New_York", "2024-03-10T02:30:00"])
        .assert()
        .failure()
        .stderr(predicate::str::contains("nonexistent datetime"));
}

#[test]
fn fails_on_ambiguous_dst_datetime() {
    // 2024-11-03 01:30 is ambiguous in America/New_York (fall back)
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt", "--from", "America/New_York", "2024-11-03T01:30:00"])
        .assert()
        .failure()
        .stderr(predicate::str::contains("ambiguous datetime"));
}
