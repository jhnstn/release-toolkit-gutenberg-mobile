# Making a release

The `bundle` directory contains the production version of the project's Javascript. This is what the WordPress apps use to avoid having to build Gutenberg.

You can rebuild those files at any time by running

```
npm run bundle
```

This is useful in case you want to use an unreleased version of the bundle in the apps. For instance, on a PR that's a work in progress, you might want to include to a specific gutenberg-mobile branch in the apps with an updated bundle so reviewers can see the latest changes before approving them (and releasing a new version).

# Release Checklist Template

Just copy this checklist and replace all occurrences of `X.XX.X` with the applicable release number, when we are ready to
cut a new release.

```
<!-- wp:heading {"level":1} -->
<h1>Gutenberg Mobile X.XX.X – Release Scenario</h1>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>This checklist is based on the <a href="https://github.com/wordpress-mobile/gutenberg-mobile/blob/develop/docs/Releasing.md#release-checklist">Release Checklist Template</a>. If you need a checklist for a new gutenberg-mobile release, please copy from that template.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>+mobilegutenberg</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>Day 1 - create the release branch, update the version</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o Visit all opened PR's in gutenberg-mobile repo that are assigned to milestone X.XX.X and leave a message with options to (i) merge the PR as soon as possible, (ii) bump the PR to the next milestone, or (iii) remove the milestone from the PR.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Check that <code>RNTAztecView.podspec</code> and <code>gutenberg/packages/react-native-aztec/RNTAztecView.podspec</code> refer to the same Aztec version.
</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Clone release scripts from <code>https://github.com/wordpress-mobile/release-toolkit-gutenberg-mobile</code> or pull the latest version if you already have it.
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Run the release script in release-toolkit-gutenberg-mobile: <code>./release_automation.sh</code>. This will take care of creating the gutenberg and gutenberg-mobile release PRs as well as WPAndroid and WPiOS integration PRs.</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>New Aztec Release</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o Make sure there is no pending Aztec PR required for this Gutenberg release. Check the commit hash referred in the gutenberg repo is in the Aztec <code>develop</code> branch. If it's not, make sure pending PRs are merged before next steps.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Open a PR on Aztec repo to update the <code>CHANGELOG.md</code> and <code>README.md</code> files with the new version name.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Create a new release and name it with the tag name from step 1. For Aztec-iOS, follow <a href="https://github.com/wordpress-mobile/AztecEditor-iOS/blob/develop/Documentation/ReleaseProcess.md">this process</a>. For Aztec-Android, releases are created via the <a href="https://github.com/wordpress-mobile/AztecEditor-Android/releases">GitHub releases page</a> by hitting the “Draft new release” button, put the tag name to be created in the tag version field and release title field, and also add the changelog to the release description. The binary assets (.zip, tar.gz files) are attached automatically after hitting “Publish release”.</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>(Optional) Specific tasks after a PR has been merged after the freeze</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o After a merge happened in gutenberg-mobile <code>release/X.XX.X</code> or in gutenberg <code>rnmobile/release-X.XX.X</code>, make sure the <code>gutenberg</code> submodule points to the right hash (and make sure the <code>rnmobile/release-X.XX.X</code> in the gutenberg repo branch has been updated)</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o If there were changes in Gutenberg repo, make sure to cherry-pick the changes that landed in the master branch back to the release branch and don't forget to run <code>npm run bundle</code> in gutenberg-mobile again if necessary.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Add the new change to the "Extra PRs that Landed After the Release Was Cut" section of the gb-mobile PR description.</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>Last Day</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o Make sure that the bundle files on the Gutenberg-Mobile release branch have been updated to include any changes to the release branch.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Merge the Gutenberg-Mobile PR to main. WARNING: Don’t merge the Gutenberg PR to master at this point.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Tag the head of Gutenberg release branch that the Gutenberg-Mobile release branch is pointing to with the <code>rnmobile/X.XX.X</code> tag.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Create a new GitHub release pointing to the tag: https://github.com/wordpress-mobile/gutenberg-mobile/releases/new?tag=vX.XX.X&target=trunk&title=Release%20X.XX.X. Include a list of changes in the release's description</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o In WPiOS update the reference to point to the <em>tag</em>. For iOS do not forget to remove <code>develop</code> branch reference near 3rd party pod specs if any.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o In WPAndroid, update the submodule to point to the merge commit on GB-Mobile <code>trunk</code>.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Main apps PRs should be ready to merge to their develop now. Merge them or get them merged.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Once everything is merged, ping our friends in #platform9 and let them know we’ve merged our release so everything is right from our side to cut the main app releases.</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>Bringing release changes back to development branches</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o If there are any conflicts in the Gutenberg PR, merge <code>master</code> into it and resolve them.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Check if you can open a PR from <code>trunk</code> to <code>develop</code> in Gutenberg Mobile without any conflicts: https://github.com/wordpress-mobile/gutenberg-mobile/compare/develop...trunk. If there are any conflicts, create a branch from <code>trunk</code> with a name like <code>merge_release_x.xx.x_to_develop</code>, merge <code>develop</code> into it, resolve any conflicts.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Open a PR from Gutenberg Mobile <code>trunk</code> (or <code>merge_release_x.xx.x_to_develop</code> branch) to <code>develop</code>.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Merge the Gutenberg PR to <code>master</code> and Gutenberg Mobile PR to <code>develop</code>.</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>AFTER the main apps have cut their release branches</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o Update the <code>gutenberg/after_X.XX.X</code> branches and open a PR against <code>develop</code>. If the branches are empty we’ll just delete them. The PR can actually get created as soon as something gets merged to the after_X.XX.X branches.&nbsp; Merge the <code>gutenberg/after_X.XX.X</code> PR(s) only AFTER the main apps have cut their release branches.</p>
<!-- /wp:paragraph -->

<!-- wp:heading {"level":3} -->
<h3>You're done</h3>
<!-- /wp:heading -->

<!-- wp:paragraph -->
<p>o Pass the baton. Ping the dev who is responsible for the next release</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>o Celebrate!</p>
<!-- /wp:paragraph -->
```
