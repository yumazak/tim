use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn converts_hours_from_stdin() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["h"])
        .write_stdin("0\n9\n23\n")
        .assert()
        .success()
        .stdout("15\n0\n14\n");
}

#[test]
fn converts_datetimes_from_stdin_and_continues_after_error() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt"])
        .write_stdin("2024-01-15T09:00:00\nnot-a-date\n2024-01-15T12:00:00\n")
        .assert()
        .code(1)
        .stdout("2024-01-15T00:00:00+00:00\n2024-01-15T03:00:00+00:00\n")
        .stderr(predicate::str::contains("invalid datetime"));
}
