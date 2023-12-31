# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class MiniReminder < Formula
  desc "Todo app for terminal"
  homepage "https://github.com/canertuzunar/miniReminder"
  version "1.0.2"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/canertuzunar/miniReminder/releases/download/v1.0.2/miniReminder_1.0.2_darwin-amd64.tar.gz"
      sha256 "bf5157658c77437ae511d2eca395f887e620864b0754064fe2020f4c0fec4884"

      def install
        bin.install "main"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/canertuzunar/miniReminder/releases/download/v1.0.2/miniReminder_1.0.2_darwin-arm64.tar.gz"
      sha256 "6f2d700c8cddffaa5d0370d0158fa2113709a44e34b9487788bb7c26536b21d9"

      def install
        bin.install "main"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/canertuzunar/miniReminder/releases/download/v1.0.2/miniReminder_1.0.2_linux-arm64.tar.gz"
      sha256 "828b1a4c17849be12b8801b029a03db9b228ee4559193eb7b48b287f39f1a3c8"

      def install
        bin.install "main"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/canertuzunar/miniReminder/releases/download/v1.0.2/miniReminder_1.0.2_linux-amd64.tar.gz"
      sha256 "9ac4314b92a2ee439b94f3eff06a5b06791819e2cb5c0c31bc6dba5188628819"

      def install
        bin.install "main"
      end
    end
  end
end
