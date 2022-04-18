// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../ZetaMPI.base.sol";
import "../ZetaInterfaces.sol";

interface ZetaToken is IERC20 {
    function burnFrom(address account, uint256 amount) external;

    function mint(
        address mintee,
        uint value,
        bytes32 internalSendHash
    ) external;
}

contract ZetaMPINonEth is ZetaMPIBase {
    constructor(
        address _zetaTokenAddress,
        address _tssAddress,
        address _tssAddressUpdater
    ) ZetaMPIBase(_zetaTokenAddress, _tssAddress, _tssAddressUpdater) {}

    function getLockedAmount() public view returns (uint) {
        return ZetaToken(zetaToken).balanceOf(address(this));
    }

    function send(ZetaInterfaces.SendInput calldata input)
        external
        override
        whenNotPaused
    {
        ZetaToken(zetaToken).burnFrom(msg.sender, input.zetaAmount);

        emit ZetaSent(
            msg.sender,
            input.destinationChainId,
            input.destinationAddress,
            input.zetaAmount,
            input.gasLimit,
            input.message,
            input.zetaParams
        );
    }

    function onReceive(
        bytes calldata originSenderAddress,
        uint256 originChainId,
        address destinationAddress,
        uint zetaAmount,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override whenNotPaused onlyTssAddress {
        ZetaToken(zetaToken).mint(
            destinationAddress,
            zetaAmount,
            internalSendHash
        );

        if (message.length > 0) {
            ZetaReceiver(destinationAddress).onZetaMessage(
                ZetaInterfaces.ZetaMessage(
                    originSenderAddress,
                    originChainId,
                    destinationAddress,
                    zetaAmount,
                    message
                )
            );
        }

        emit ZetaReceived(
            originSenderAddress,
            originChainId,
            destinationAddress,
            zetaAmount,
            message,
            internalSendHash
        );
    }

    function onRevert(
        address originSenderAddress,
        uint256 originChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 zetaAmount,
        bytes calldata message,
        bytes32 internalSendHash
    ) external override whenNotPaused onlyTssAddress {
        ZetaToken(zetaToken).mint(
            originSenderAddress,
            zetaAmount,
            internalSendHash
        );

        if (message.length > 0) {
            ZetaReceiver(originSenderAddress).onZetaRevert(
                ZetaInterfaces.ZetaRevert(
                    originSenderAddress,
                    originChainId,
                    destinationAddress,
                    destinationChainId,
                    zetaAmount,
                    message
                )
            );
        }

        emit ZetaReverted(
            originSenderAddress,
            originChainId,
            destinationChainId,
            destinationAddress,
            zetaAmount,
            message,
            internalSendHash
        );
    }
}
