{-# LANGUAGE OverloadedStrings #-}
module Lib
    ( someFunc
    ) where

import qualified Data.Text.IO as T
import Acme.Missiles

someFunc :: IO ()
-- someFunc = T.putStrLn "someFunc"
someFunc = launchMissiles
