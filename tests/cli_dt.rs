use assert_cmd::Command;

#[test]
fn converts_datetime_with_defaults() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt", "2024-01-15T09:00:00"])
        .assert()
        .success()
        .stdout("2024-01-15T00:00:00+00:00\n");
}

#[test]
fn converts_datetime_with_space_separator() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt", "2024-01-15 09:00:00"])
        .assert()
        .success()
        .stdout("2024-01-15T00:00:00+00:00\n");
}

#[test]
fn uses_embedded_offset_when_present() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt", "2024-01-15T09:00:00+09:00"])
        .assert()
        .success()
        .stdout("2024-01-15T00:00:00+00:00\n");
}

#[test]
fn converts_datetime_without_seconds() {
    Command::cargo_bin("tim")
        .unwrap()
        .args(["dt", "2024-01-15T09:00"])
        .assert()
        .success()
        .stdout("2024-01-15T00:00:00+00:00\n");
}
